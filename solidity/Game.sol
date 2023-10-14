// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";

contract GameInfoContract is Ownable {
    struct GameInfo {
        int64 DomainSeparationTag;
        string ReplayCID;
        bytes Entropy;
        uint64 VRFHeight;
        bytes VRFProof;
        bytes GameResults;
    }

    mapping(string => GameInfo) gameInfoMap;

    function setGameInfo(
        string memory _roundId,
        GameInfo memory _info
    ) public onlyOwner {
        gameInfoMap[_roundId] = _info;
    }

    function getGameInfo(string memory _roundId) public view returns (GameInfo memory) {
        GameInfo memory info = gameInfoMap[_roundId];
        require(bytes(info.ReplayCID).length > 0, "Game info not found");
        return info;
    }
}