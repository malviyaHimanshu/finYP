// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract Audit {
    //a simple unsigned int variable
    uint private storedData;
    bool private auditPass;
    bytes storedParams;
    bytes storedG;
    bytes storedU;
    bytes storedPKO;
    bytes storedPKS;
    string storedSign;

    /*struct PBCvars { // Struct
        bytes storedParams;
        bytes storedG;
        bytes storedU;
    }

    struct PKs {
        bytes storedPKO;
        bytes storedPKS;
    }*/

    struct ownerKeys {
        bytes storedPKT;
        bytes storedSKH;
    }

    ownerKeys public keys;

    ownerKeys[] public keyArr;
    
    //func to set the value of the var
    function set (uint x) public {
        storedData = x;
    }

    function setParamsGU (bytes memory input1, bytes memory input2, bytes memory input3) public {
        storedParams = input1;
        storedG = input2;
        storedU = input3;
    }
    
    function setPKO (bytes memory input) public {
        storedPKO = input;
    }
    
    function setPKS (bytes memory input) public {
        storedPKS = input;
    }

    function setSign (string memory input) public {
        storedSign = input;
    }

    function setOwnerKeys (bytes memory input1, bytes memory input2) public {
        keys = ownerKeys(input1, input2);
        keyArr.push(keys);
    }

    //func to get the value of the var
    function get() public view returns (uint retVal) {
        return storedData;
    }

    function getParamsGU() public view returns (bytes memory, bytes memory, bytes memory) {
        return (storedParams, storedG, storedU);
    }

    function getPKT() public view returns (bytes memory) {
        return keys.storedPKT;
    }

    function getPKS() public view returns (bytes memory) {
        return storedPKS;
    }

    function getSKH() public view returns (bytes memory) {
        return keys.storedSKH;
    }

    
    constructor(){
        auditPass = false;
    }
        
    //func to test newly written pre-compile
    function sendAudit(bytes memory input) public {
        bytes[1] memory output;
        bool ret;
        uint256 len = input.length + 32;
        
        assembly {
            ret := iszero(call(not(0), 0x09, 0, input, len, output, 0x01))
        }
        assert(!ret);
        
        if (keccak256(output[0])==keccak256("1")) {
            auditPass = true;
        }
    }
    
    function queryAudit() public view returns (bool){
        return auditPass;
    }

}

