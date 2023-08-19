#include "card.h"
#include <iostream>

Card::Card() {}

Card::Card(int rank, int suit) {
    this->rank = rank;
    this->suit = suit;
}

bool Card::Bigger(Card card) {
    if (this->rank < card.rank) {
        return false;
    }

    if (this->rank == card.rank && this->suit < card.suit) {
        return false;
    }

    return true;
}

std::string Card::String() {

    std::string rank = "";
    std::string suit = "";

    switch(this->rank) {
        case RANK2:
            rank = "2";
            break;
        case RANK3:
            rank = "3";
            break;
        case RANK4:
            rank = "4";
            break;
        case RANK5:
            rank = "5";
            break;
        case RANK6:
            rank = "6";
            break;
        case RANK7:
            rank = "7";
            break;
        case RANK8:
            rank = "8";
            break;
        case RANK9:
            rank = "9";
            break;
        case RANK10:
            rank = "10";
            break;
        case RANKJ:
            rank = "J";
            break;
        case RANKQ:
            rank = "Q";
            break;
        case RANKK:
            rank = "K";
            break;
        case RANKA:
            rank = "A";
            break;
    }

    switch (this->suit) {
        case SuitClub:
            suit = "Club";
            break;
        case SuitDiamond:
            suit = "Diamond";
            break;
        case SuitHeart:
            suit = "Heart";
            break;
        case SuitSpade:
            suit = "Spade";
            break;
    }

    return suit + rank;
}