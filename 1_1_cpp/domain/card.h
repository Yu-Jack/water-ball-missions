#include <string>

#ifndef CARD_H
#define CARD_H

enum Rank {
    RANK2 = 2,
    RANK3,
	RANK4,
	RANK5,
	RANK6,
	RANK7,
	RANK8,
	RANK9,
	RANK10,
	RANKJ,
	RANKQ,
	RANKK,
	RANKA
};

enum Suit {
	CLUB = 1,
	DIAMOND,
	HEART,
	SPADE
};

class Card {
private:
    Rank rank;
    Suit suit;

public:
    Card(int rank, int suit);
    bool Bigger(Card card);
    std::string String();
};

#endif