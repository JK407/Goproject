pragma solidity ^0.6.0;


contract UserManager {
    address payable public owner;
    
    mapping(uint8 => string) accounts;    //id -> passwd
    mapping(uint8 => address) ips;       //IP address
    mapping(address => bool) isExist;
    
    event Login(uint8 id, uint time);
    event Register();
    event SetPassword();
    
    constructor () public {
        owner = msg.sender;
    }
    event changePassword(string  password,string  newpassword);
    function login(uint8 id, string memory password) public returns (bool) {
        require(ips[id] == msg.sender);
        if (keccak256(abi.encodePacked(id)) == keccak256(abi.encodePacked(password))) {
            emit Login(id, now);
            return true; 
        }
        return false;
    }
    
    function getIP(uint8 id) public view returns (address) {
        require(ips[id] != address(0));
        return ips[id];
    }
    
    function register(uint8 id,string memory password) public returns (bool) {
        require(!isExist[msg.sender],"");
        require(keccak256(abi.encodePacked(id)) == keccak256(abi.encodePacked(password)),"");
        accounts[id] = password;
        isExist[msg.sender] = true;
        return true;
    }
    
    function setPassword(string memory password,string memory newpassword) public returns (bool) {
        emit changePassword(password,newpassword);
        password=newpassword;
        
    }
  }