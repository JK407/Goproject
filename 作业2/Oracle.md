pragma solidity ^0.6.0;

/**
* @title CheckpointOracle
* @author Gary Rong<garyrong@ethereum.org>, Martin Swende <martin.swende@ethereum.org>
* @dev Implementation of the blockchain checkpoint registrar.
  */
  contract CheckpointOracle {
  /*
  事件
  */

  // NewCheckpointVote is emitted when a new checkpoint proposal receives a vote.
  event NewCheckpointVote(uint64 indexed index, bytes32 checkpointHash, uint8 v, bytes32 r, bytes32 s);

  /*
  公共函数
  */
  constructor(address[] memory _adminlist, uint _sectionSize, uint _processConfirms, uint _threshold) public {
  for (uint i = 0; i < _adminlist.length; i++) {
  admins[_adminlist[i]] = true;
  adminList.push(_adminlist[i]);
  }
  sectionSize = _sectionSize;
  processConfirms = _processConfirms;
  threshold = _threshold;
  }

  /*
      dev获取最稳定的检查点信息，返回节索引，检查点哈希，和与检查点有关的返回块高度
      */
      function GetLatestCheckpoint()
      view
      public
      returns(uint64, bytes32, uint) {
      return (sectionIndex, hash, height);
      }

  // SetCheckpoint 设立一个新的检查点，用来接受签名列表
  // @_recentNumber: 一个最近的区块，用于重播保护
  // @_recentHash : recentnumber的hash
  // @_hash : 在_sectionIndex处设置的哈希
  // @_sectionIndex : 要设置的部分索引
  // @v : v值列表
  // @r : r值列表
  // @s : s值列表
  function SetCheckpoint(
  uint _recentNumber,
  bytes32 _recentHash,
  bytes32 _hash,
  uint64 _sectionIndex,
  uint8[] memory v,
  bytes32[] memory r,
  bytes32[] memory s)
  public
  returns (bool)
  {
  // 确保发件人被授权。
  require(admins[msg.sender]);

       // 这些检查会重播保护，因此无法在fork上重播，
       // 意外或故意
       require(blockhash(_recentNumber) == _recentHash);

       // 确保签名批次有效。
       require(v.length == r.length);
       require(v.length == s.length);

       // 过滤掉“未来”检查点。
       if (block.number < (_sectionIndex+1)*sectionSize+processConfirms) {
           return false;
       }
       // 过滤掉“旧”公告
       if (_sectionIndex < sectionIndex) {
           return false;
       }
       // 过滤掉“过时”的公告
       if (_sectionIndex == sectionIndex && (_sectionIndex != 0 || height != 0)) {
           return false;
       }
       //过滤掉“无效”公告
       if (_hash == ""){
           return false;
       }

       // EIP 191样式签名
       //
       // 计算哈希值以进行验证时的参数
       // 1: byte(0x19) - 初始0x19字节
       // 2: byte(0) - 版本字节（带有预期验证器的数据）
       // 3: this - 验证者地址
       // --  特定于应用的数据
       // 4 : checkpoint section_index(uint64) -----检查点section_index（uint64
       // 5 : checkpoint hash (bytes32)-----检查点哈希（bytes32）
       //     hash = keccak256(checkpoint_index, section_head, cht_root, bloom_root)
       bytes32 signedHash = keccak256(abi.encodePacked(byte(0x19), byte(0), this, _sectionIndex, _hash));

       address lastVoter = address(0);

       // 为了使我们不必维护谁已经映射
       // 投票，并且我们不想重复投票，签名必须
       // 严格按顺序提交。
       for (uint idx = 0; idx < v.length; idx++){
           address signer = ecrecover(signedHash, v[idx], r[idx], s[idx]);
           require(admins[signer]);
           require(uint256(signer) > uint256(lastVoter));
           lastVoter = signer;
           emit NewCheckpointVote(_sectionIndex, _hash, v[idx], r[idx], s[idx]);

           // 存在足够的签名，更新最新的检查点。
           if (idx+1 >= threshold){
               hash = _hash;
               height = block.number;
               sectionIndex = _sectionIndex;
               return true;
           }
       }
       // 我们不应该在这里结束，恢复未发出的事件
       revert();
  }

  /**
    * @dev获取所有管理员地址
    * @返回地址列表
      */
      function GetAllAdmin()
      public
      view
      returns(address[] memory)
      {
      address[] memory ret = new address[](adminList.length);
      for (uint i = 0; i < adminList.length; i++) {
      ret[i] = adminList[i];
      }
      return ret;
      }

  /*
  Fields
  */
  // A map 有权更新CHT和Bloom Trie root的管理员用户

  //管理员用户列表，以便我们可以获得所有管理员用户。
  // 最新存储的部分ID
  uint64 sectionIndex;

  // 与最新注册的检查点关联的块高度。

  // 最近注册的检查点的哈希值。

  // 创建检查点的频率
  //
  // 默认值应与以太坊中的检查点大小（32768）相同。
  // 可以注册检查点之前所需的确认数。
  // 我们必须确保注册的检查点不会由于以下原因而无效
  //链重组。
  //
  // 默认值应与检查点过程确认相同（256）
  // 在以太坊中。
  uint processConfirms;

  // 完成稳定的检查点所需的签名。
  uint threshold;
  }