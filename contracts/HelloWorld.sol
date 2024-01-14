// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0 <0.9.0;

contract HelloWorld {
    string public data;

    event Received(address sender, uint amount);

    receive() external payable {
        emit Received(msg.sender, msg.value);
    }

    function Hello() public pure returns (string memory) {
        return "Hello World";
    }

    function Greet(string memory str) public pure returns (string memory) {
        require(bytes(str).length > 0, "string must not be empty");
        return str;
    }

    function SetData(string memory str) public {
        require(bytes(str).length > 0, "string must not be empty");
        data = str;
    }

    function GetData() public view returns (string memory) {
        return data;
    }

    function GetBalance() public view returns (uint256) {
        return address(this).balance;
    }

    function Withdraw() public {
        payable(msg.sender).transfer(address(this).balance);
    }
}