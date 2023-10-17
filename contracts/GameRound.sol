// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract GameRound {
    struct Info {
        string GameID;
        string RoundID;
        string ReplayID;
        string PlayerIDs;
    }

    struct PlayerResult {
        string PlayerID;
        uint64 CurrentScore;
        uint64 WinScore;

    }

    struct Result {
        PlayerResult []PlayerResults;
    }

    struct Replay {
        int64 DomainSeparationTag;
        string Address;
        Info GameInfo;
        uint64 VRFHeight;
        string HashFunc;
        string Seed;
        Result GameResult;
    }
}
