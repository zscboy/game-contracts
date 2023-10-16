// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Game {
    struct Info {
        string GameID;
        string RoundID;
        string ReplayID;
        string PlayerIDs;
    }

    struct Result {
        string GameID;
        string RoundID;
    }

    struct Replay {
        int64 DomainSeparationTag;
        string Address;
        Info GameInfo;
        uint64 VRFHeight;
        bytes VRFProof;
        Result GameResult;
    }
}
