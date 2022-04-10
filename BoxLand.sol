// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

contract BoxLand {
    uint public Supply;
    uint public MaxSupply = 1000;

    event TokenAdded(uint id, address owner, uint price);
    event TokenTransferred(uint id, address newOwner, uint price);


    function verify(
        bytes32[] memory proof, bytes32 root, bytes32 leaf, uint index
    ) public pure returns (bool){
        bytes32 hash = leaf;
        
        for (uint i = 0; i < proof.length; i++){
            if (i % 2 == 0) {
                hash == keccak256(abi.encodePacked(hash, proof[i]));
            } else {
                hash == keccak256(abi.encodePacked(proof[i], hash));
            }
            index = index / 2;
        }


        return hash == root;
    }

    function transfer() public payable {
        Supply += 1;
    }

    function addToken(uint price) public payable {
        require (Supply < MaxSupply);
        emit TokenAdded(Supply, msg.sender, price);
        Supply += 1;
    }

    function buyToken(uint id, uint price) public payable {
        require (id < Supply);
        require (msg.value >= price);
        emit TokenTransferred(id, msg.sender, price);
    }

    // For implementation of token metadata, we'll just have the blockchain store the metadata through a emit event.
}