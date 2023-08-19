#include <string>

#ifndef CARD_H
#define CARD_H

class Card {
private:
    int rank;
    int suit;

public:
    static const int RANK2 = 1;
    static const int RANK3 = 2;
    static const int RANK4 = 3;
    static const int RANK5 = 4;
    static const int RANK6 = 5;
    static const int RANK7 = 6;
    static const int RANK8 = 7;
    static const int RANK9 = 8;
    static const int RANK10 = 9;
    static const int RANKJ = 10;
    static const int RANKQ = 11;
    static const int RANKK = 12;
    static const int RANKA = 13;

    static const int SuitClub = 1;
    static const int SuitDiamond = 2;
    static const int SuitHeart = 3;
    static const int SuitSpade = 4;

    Card();
    Card(int rank, int suit);
    bool Bigger(Card card);
    std::string String();
};

#endif