pragma solidity ^0.4.4;


import "../contracts_zeppelin/token/StandardToken.sol";
import "../contracts_zeppelin/ownership/Ownable.sol";


/*
 *  Sonm Dummy Token for test
*
 *  
 *
 * Very simple ERC20 Token example, where all tokens are pre-assigned
 * to the creator. Note they can later distribute these tokens
 * as they wish using `transfer` and other `StandardToken` functions.
 *
 *  simple standard
 */
contract SonmDummyToken is StandardToken, Ownable {

    string public name = "Sonm Dummy Token";

    string public symbol = "SDT";

    uint public decimals = 18;

    uint public INITIAL_SUPPLY = 1000000 * 1 ether;

    function SonmDummyToken() {
        totalSupply = INITIAL_SUPPLY;
        balances[msg.sender] = INITIAL_SUPPLY;
    }

}
