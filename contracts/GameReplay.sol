// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Ownable.sol";
import "./GameRound.sol";

contract GameReplayContract is Ownable {
    mapping(string => GameRound.Replay) gameReplayMap;

    function saveGameReplay(GameRound.Replay[] memory _replays) public onlyOwner {
        require(_replays.length > 0, "_replays can not empty");

        for (uint i = 0; i < _replays.length; i++) {
             gameReplayMap[_replays[i].GameInfo.RoundID] = _replays[i];
        }
    }

    function getGameReplay(string memory _roundID) public view returns (GameRound.Replay memory) {
        GameRound.Replay memory replay = gameReplayMap[_roundID];
        require(bytes(replay.Address).length > 0,  string(abi.encodePacked("Game replay not found: ", _roundID)));
        
        return replay;
    }
}