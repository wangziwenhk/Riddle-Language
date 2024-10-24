# Generated from C:/Users/wangz/PycharmProjects/Riddle-Language/RiddleParser.g4 by ANTLR 4.13.2
# encoding: utf-8
from antlr4 import *
from io import StringIO
import sys
if sys.version_info[1] > 5:
	from typing import TextIO
else:
	from typing.io import TextIO

def serializedATN():
    return [
        4,1,66,487,2,0,7,0,2,1,7,1,2,2,7,2,2,3,7,3,2,4,7,4,2,5,7,5,2,6,7,
        6,2,7,7,7,2,8,7,8,2,9,7,9,2,10,7,10,2,11,7,11,2,12,7,12,2,13,7,13,
        2,14,7,14,2,15,7,15,2,16,7,16,2,17,7,17,2,18,7,18,2,19,7,19,2,20,
        7,20,2,21,7,21,2,22,7,22,2,23,7,23,2,24,7,24,2,25,7,25,2,26,7,26,
        2,27,7,27,2,28,7,28,2,29,7,29,1,0,5,0,62,8,0,10,0,12,0,65,9,0,1,
        0,3,0,68,8,0,1,1,1,1,1,2,1,2,3,2,74,8,2,1,2,3,2,77,8,2,1,2,1,2,3,
        2,81,8,2,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,1,3,5,3,
        96,8,3,10,3,12,3,99,9,3,1,3,3,3,102,8,3,1,4,5,4,105,8,4,10,4,12,
        4,108,9,4,1,5,1,5,1,5,1,6,1,6,1,6,1,7,1,7,1,7,1,7,1,7,1,7,1,7,1,
        7,1,7,1,7,1,7,1,7,1,7,1,7,1,7,3,7,131,8,7,1,8,1,8,1,8,5,8,136,8,
        8,10,8,12,8,139,9,8,1,8,3,8,142,8,8,1,9,1,9,1,9,1,9,1,9,5,9,149,
        8,9,10,9,12,9,152,9,9,1,9,1,9,1,9,3,9,157,8,9,1,10,1,10,1,10,1,10,
        1,10,1,10,1,10,1,10,3,10,167,8,10,1,10,1,10,1,10,1,10,1,11,1,11,
        1,11,3,11,176,8,11,1,11,1,11,3,11,180,8,11,1,11,1,11,3,11,184,8,
        11,1,11,1,11,1,11,1,12,1,12,1,12,1,12,1,12,1,12,1,13,1,13,1,13,1,
        13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,1,13,3,
        13,211,8,13,1,14,1,14,3,14,215,8,14,1,15,1,15,1,15,1,15,1,15,1,15,
        1,16,1,16,1,16,1,16,1,16,3,16,228,8,16,1,16,1,16,1,17,1,17,1,17,
        1,17,1,17,1,18,1,18,1,18,1,18,1,18,1,18,1,18,3,18,244,8,18,1,18,
        1,18,1,18,5,18,249,8,18,10,18,12,18,252,9,18,1,19,1,19,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,3,20,357,8,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,
        1,20,1,20,1,20,1,20,1,20,1,20,1,20,1,20,5,20,421,8,20,10,20,12,20,
        424,9,20,1,21,1,21,1,22,1,22,3,22,430,8,22,1,23,1,23,1,23,1,23,3,
        23,436,8,23,1,24,1,24,1,25,1,25,1,25,1,26,1,26,1,26,1,26,1,26,1,
        26,1,26,1,26,3,26,451,8,26,1,27,1,27,3,27,455,8,27,1,28,1,28,1,28,
        5,28,460,8,28,10,28,12,28,463,9,28,1,28,3,28,466,8,28,1,29,1,29,
        1,29,1,29,1,29,1,29,1,29,3,29,475,8,29,1,29,1,29,1,29,1,29,1,29,
        5,29,482,8,29,10,29,12,29,485,9,29,1,29,0,3,36,40,58,30,0,2,4,6,
        8,10,12,14,16,18,20,22,24,26,28,30,32,34,36,38,40,42,44,46,48,50,
        52,54,56,58,0,1,2,0,30,30,51,51,545,0,67,1,0,0,0,2,69,1,0,0,0,4,
        80,1,0,0,0,6,101,1,0,0,0,8,106,1,0,0,0,10,109,1,0,0,0,12,112,1,0,
        0,0,14,130,1,0,0,0,16,141,1,0,0,0,18,156,1,0,0,0,20,158,1,0,0,0,
        22,172,1,0,0,0,24,188,1,0,0,0,26,210,1,0,0,0,28,212,1,0,0,0,30,216,
        1,0,0,0,32,222,1,0,0,0,34,231,1,0,0,0,36,243,1,0,0,0,38,253,1,0,
        0,0,40,356,1,0,0,0,42,425,1,0,0,0,44,429,1,0,0,0,46,435,1,0,0,0,
        48,437,1,0,0,0,50,439,1,0,0,0,52,450,1,0,0,0,54,454,1,0,0,0,56,465,
        1,0,0,0,58,474,1,0,0,0,60,62,3,4,2,0,61,60,1,0,0,0,62,65,1,0,0,0,
        63,61,1,0,0,0,63,64,1,0,0,0,64,68,1,0,0,0,65,63,1,0,0,0,66,68,5,
        0,0,1,67,63,1,0,0,0,67,66,1,0,0,0,68,1,1,0,0,0,69,70,7,0,0,0,70,
        3,1,0,0,0,71,73,3,6,3,0,72,74,5,30,0,0,73,72,1,0,0,0,73,74,1,0,0,
        0,74,76,1,0,0,0,75,77,5,51,0,0,76,75,1,0,0,0,76,77,1,0,0,0,77,81,
        1,0,0,0,78,81,5,30,0,0,79,81,5,51,0,0,80,71,1,0,0,0,80,78,1,0,0,
        0,80,79,1,0,0,0,81,5,1,0,0,0,82,102,3,10,5,0,83,102,3,12,6,0,84,
        102,3,30,15,0,85,102,3,20,10,0,86,102,3,14,7,0,87,102,3,22,11,0,
        88,102,3,24,12,0,89,102,3,26,13,0,90,102,3,28,14,0,91,102,3,32,16,
        0,92,102,3,40,20,0,93,97,5,27,0,0,94,96,3,4,2,0,95,94,1,0,0,0,96,
        99,1,0,0,0,97,95,1,0,0,0,97,98,1,0,0,0,98,100,1,0,0,0,99,97,1,0,
        0,0,100,102,5,28,0,0,101,82,1,0,0,0,101,83,1,0,0,0,101,84,1,0,0,
        0,101,85,1,0,0,0,101,86,1,0,0,0,101,87,1,0,0,0,101,88,1,0,0,0,101,
        89,1,0,0,0,101,90,1,0,0,0,101,91,1,0,0,0,101,92,1,0,0,0,101,93,1,
        0,0,0,102,7,1,0,0,0,103,105,3,4,2,0,104,103,1,0,0,0,105,108,1,0,
        0,0,106,104,1,0,0,0,106,107,1,0,0,0,107,9,1,0,0,0,108,106,1,0,0,
        0,109,110,5,10,0,0,110,111,3,42,21,0,111,11,1,0,0,0,112,113,5,9,
        0,0,113,114,3,42,21,0,114,13,1,0,0,0,115,116,5,1,0,0,116,117,5,52,
        0,0,117,118,5,29,0,0,118,131,3,58,29,0,119,120,5,1,0,0,120,121,5,
        52,0,0,121,122,5,33,0,0,122,131,3,40,20,0,123,124,5,1,0,0,124,125,
        5,52,0,0,125,126,5,29,0,0,126,127,3,58,29,0,127,128,5,33,0,0,128,
        129,3,40,20,0,129,131,1,0,0,0,130,115,1,0,0,0,130,119,1,0,0,0,130,
        123,1,0,0,0,131,15,1,0,0,0,132,133,3,40,20,0,133,134,5,31,0,0,134,
        136,1,0,0,0,135,132,1,0,0,0,136,139,1,0,0,0,137,135,1,0,0,0,137,
        138,1,0,0,0,138,140,1,0,0,0,139,137,1,0,0,0,140,142,3,40,20,0,141,
        137,1,0,0,0,141,142,1,0,0,0,142,17,1,0,0,0,143,144,5,52,0,0,144,
        145,5,29,0,0,145,146,3,58,29,0,146,147,5,31,0,0,147,149,1,0,0,0,
        148,143,1,0,0,0,149,152,1,0,0,0,150,148,1,0,0,0,150,151,1,0,0,0,
        151,153,1,0,0,0,152,150,1,0,0,0,153,154,5,52,0,0,154,155,5,29,0,
        0,155,157,3,58,29,0,156,150,1,0,0,0,156,157,1,0,0,0,157,19,1,0,0,
        0,158,159,5,7,0,0,159,160,5,52,0,0,160,161,5,23,0,0,161,162,3,18,
        9,0,162,166,5,24,0,0,163,164,5,40,0,0,164,165,5,34,0,0,165,167,3,
        58,29,0,166,163,1,0,0,0,166,167,1,0,0,0,167,168,1,0,0,0,168,169,
        5,27,0,0,169,170,3,8,4,0,170,171,5,28,0,0,171,21,1,0,0,0,172,173,
        5,3,0,0,173,175,5,23,0,0,174,176,3,14,7,0,175,174,1,0,0,0,175,176,
        1,0,0,0,176,177,1,0,0,0,177,179,5,30,0,0,178,180,3,40,20,0,179,178,
        1,0,0,0,179,180,1,0,0,0,180,181,1,0,0,0,181,183,5,30,0,0,182,184,
        3,6,3,0,183,182,1,0,0,0,183,184,1,0,0,0,184,185,1,0,0,0,185,186,
        5,24,0,0,186,187,3,4,2,0,187,23,1,0,0,0,188,189,5,4,0,0,189,190,
        5,23,0,0,190,191,3,40,20,0,191,192,5,24,0,0,192,193,3,4,2,0,193,
        25,1,0,0,0,194,195,5,5,0,0,195,196,5,23,0,0,196,197,3,40,20,0,197,
        198,5,24,0,0,198,199,3,4,2,0,199,200,6,13,-1,0,200,211,1,0,0,0,201,
        202,5,5,0,0,202,203,5,23,0,0,203,204,3,40,20,0,204,205,5,24,0,0,
        205,206,3,4,2,0,206,207,5,6,0,0,207,208,3,4,2,0,208,209,6,13,-1,
        0,209,211,1,0,0,0,210,194,1,0,0,0,210,201,1,0,0,0,211,27,1,0,0,0,
        212,214,5,8,0,0,213,215,3,4,2,0,214,213,1,0,0,0,214,215,1,0,0,0,
        215,29,1,0,0,0,216,217,5,11,0,0,217,218,3,42,21,0,218,219,5,27,0,
        0,219,220,3,8,4,0,220,221,5,28,0,0,221,31,1,0,0,0,222,223,5,21,0,
        0,223,224,5,27,0,0,224,225,3,8,4,0,225,227,5,28,0,0,226,228,3,2,
        1,0,227,226,1,0,0,0,227,228,1,0,0,0,228,229,1,0,0,0,229,230,3,34,
        17,0,230,33,1,0,0,0,231,232,5,22,0,0,232,233,5,23,0,0,233,234,3,
        14,7,0,234,235,5,24,0,0,235,35,1,0,0,0,236,237,6,18,-1,0,237,238,
        5,52,0,0,238,239,5,23,0,0,239,240,3,16,8,0,240,241,5,24,0,0,241,
        244,1,0,0,0,242,244,5,52,0,0,243,236,1,0,0,0,243,242,1,0,0,0,244,
        250,1,0,0,0,245,246,10,1,0,0,246,247,5,48,0,0,247,249,3,36,18,2,
        248,245,1,0,0,0,249,252,1,0,0,0,250,248,1,0,0,0,250,251,1,0,0,0,
        251,37,1,0,0,0,252,250,1,0,0,0,253,254,3,36,18,0,254,39,1,0,0,0,
        255,256,6,20,-1,0,256,257,5,35,0,0,257,258,3,58,29,0,258,259,5,34,
        0,0,259,260,5,23,0,0,260,261,3,38,19,0,261,262,5,24,0,0,262,357,
        1,0,0,0,263,264,5,23,0,0,264,265,3,40,20,0,265,266,5,24,0,0,266,
        357,1,0,0,0,267,268,5,44,0,0,268,357,3,40,20,44,269,270,5,39,0,0,
        270,357,3,40,20,43,271,272,5,40,0,0,272,357,3,40,20,42,273,274,5,
        39,0,0,274,275,5,39,0,0,275,357,3,38,19,0,276,277,3,38,19,0,277,
        278,5,39,0,0,278,279,5,39,0,0,279,357,1,0,0,0,280,281,5,40,0,0,281,
        282,5,40,0,0,282,357,3,38,19,0,283,284,3,38,19,0,284,285,5,40,0,
        0,285,286,5,40,0,0,286,357,1,0,0,0,287,357,3,36,18,0,288,289,3,38,
        19,0,289,290,5,33,0,0,290,291,3,40,20,17,291,357,1,0,0,0,292,293,
        3,38,19,0,293,294,5,39,0,0,294,295,5,33,0,0,295,296,3,40,20,16,296,
        357,1,0,0,0,297,298,3,38,19,0,298,299,5,40,0,0,299,300,5,33,0,0,
        300,301,3,40,20,15,301,357,1,0,0,0,302,303,3,38,19,0,303,304,5,41,
        0,0,304,305,5,33,0,0,305,306,3,40,20,14,306,357,1,0,0,0,307,308,
        3,38,19,0,308,309,5,42,0,0,309,310,5,33,0,0,310,311,3,40,20,13,311,
        357,1,0,0,0,312,313,3,38,19,0,313,314,5,43,0,0,314,315,5,33,0,0,
        315,316,3,40,20,12,316,357,1,0,0,0,317,318,3,38,19,0,318,319,5,39,
        0,0,319,320,5,33,0,0,320,321,3,40,20,11,321,357,1,0,0,0,322,323,
        3,38,19,0,323,324,5,45,0,0,324,325,5,33,0,0,325,326,3,40,20,10,326,
        357,1,0,0,0,327,328,3,38,19,0,328,329,5,46,0,0,329,330,5,33,0,0,
        330,331,3,40,20,9,331,357,1,0,0,0,332,333,3,38,19,0,333,334,5,47,
        0,0,334,335,5,33,0,0,335,336,3,40,20,8,336,357,1,0,0,0,337,338,3,
        38,19,0,338,339,5,36,0,0,339,340,5,33,0,0,340,341,3,40,20,7,341,
        357,1,0,0,0,342,343,3,38,19,0,343,344,5,37,0,0,344,345,5,33,0,0,
        345,346,3,40,20,6,346,357,1,0,0,0,347,348,3,38,19,0,348,349,5,38,
        0,0,349,350,5,33,0,0,350,351,3,40,20,5,351,357,1,0,0,0,352,357,3,
        48,24,0,353,357,3,44,22,0,354,357,3,46,23,0,355,357,5,20,0,0,356,
        255,1,0,0,0,356,263,1,0,0,0,356,267,1,0,0,0,356,269,1,0,0,0,356,
        271,1,0,0,0,356,273,1,0,0,0,356,276,1,0,0,0,356,280,1,0,0,0,356,
        283,1,0,0,0,356,287,1,0,0,0,356,288,1,0,0,0,356,292,1,0,0,0,356,
        297,1,0,0,0,356,302,1,0,0,0,356,307,1,0,0,0,356,312,1,0,0,0,356,
        317,1,0,0,0,356,322,1,0,0,0,356,327,1,0,0,0,356,332,1,0,0,0,356,
        337,1,0,0,0,356,342,1,0,0,0,356,347,1,0,0,0,356,352,1,0,0,0,356,
        353,1,0,0,0,356,354,1,0,0,0,356,355,1,0,0,0,357,422,1,0,0,0,358,
        359,10,36,0,0,359,360,5,41,0,0,360,421,3,40,20,37,361,362,10,35,
        0,0,362,363,5,42,0,0,363,421,3,40,20,36,364,365,10,34,0,0,365,366,
        5,43,0,0,366,421,3,40,20,35,367,368,10,33,0,0,368,369,5,39,0,0,369,
        421,3,40,20,34,370,371,10,32,0,0,371,372,5,40,0,0,372,421,3,40,20,
        33,373,374,10,31,0,0,374,375,5,36,0,0,375,421,3,40,20,32,376,377,
        10,30,0,0,377,378,5,37,0,0,378,421,3,40,20,31,379,380,10,29,0,0,
        380,381,5,38,0,0,381,421,3,40,20,30,382,383,10,28,0,0,383,384,5,
        34,0,0,384,421,3,40,20,29,385,386,10,27,0,0,386,387,5,35,0,0,387,
        421,3,40,20,28,388,389,10,26,0,0,389,390,5,34,0,0,390,391,5,33,0,
        0,391,421,3,40,20,27,392,393,10,25,0,0,393,394,5,35,0,0,394,395,
        5,33,0,0,395,421,3,40,20,26,396,397,10,24,0,0,397,398,5,32,0,0,398,
        421,3,40,20,25,399,400,10,23,0,0,400,401,5,44,0,0,401,402,5,33,0,
        0,402,421,3,40,20,24,403,404,10,22,0,0,404,405,5,45,0,0,405,421,
        3,40,20,23,406,407,10,21,0,0,407,408,5,47,0,0,408,421,3,40,20,22,
        409,410,10,20,0,0,410,411,5,46,0,0,411,421,3,40,20,21,412,413,10,
        19,0,0,413,414,5,45,0,0,414,415,5,45,0,0,415,421,3,40,20,20,416,
        417,10,18,0,0,417,418,5,46,0,0,418,419,5,46,0,0,419,421,3,40,20,
        19,420,358,1,0,0,0,420,361,1,0,0,0,420,364,1,0,0,0,420,367,1,0,0,
        0,420,370,1,0,0,0,420,373,1,0,0,0,420,376,1,0,0,0,420,379,1,0,0,
        0,420,382,1,0,0,0,420,385,1,0,0,0,420,388,1,0,0,0,420,392,1,0,0,
        0,420,396,1,0,0,0,420,399,1,0,0,0,420,403,1,0,0,0,420,406,1,0,0,
        0,420,409,1,0,0,0,420,412,1,0,0,0,420,416,1,0,0,0,421,424,1,0,0,
        0,422,420,1,0,0,0,422,423,1,0,0,0,423,41,1,0,0,0,424,422,1,0,0,0,
        425,426,5,52,0,0,426,43,1,0,0,0,427,430,3,52,26,0,428,430,3,50,25,
        0,429,427,1,0,0,0,429,428,1,0,0,0,430,45,1,0,0,0,431,432,5,16,0,
        0,432,436,6,23,-1,0,433,434,5,17,0,0,434,436,6,23,-1,0,435,431,1,
        0,0,0,435,433,1,0,0,0,436,47,1,0,0,0,437,438,5,63,0,0,438,49,1,0,
        0,0,439,440,5,57,0,0,440,441,6,25,-1,0,441,51,1,0,0,0,442,443,5,
        54,0,0,443,451,6,26,-1,0,444,445,5,53,0,0,445,451,6,26,-1,0,446,
        447,5,56,0,0,447,451,6,26,-1,0,448,449,5,55,0,0,449,451,6,26,-1,
        0,450,442,1,0,0,0,450,444,1,0,0,0,450,446,1,0,0,0,450,448,1,0,0,
        0,451,53,1,0,0,0,452,455,3,40,20,0,453,455,3,58,29,0,454,452,1,0,
        0,0,454,453,1,0,0,0,455,55,1,0,0,0,456,457,3,54,27,0,457,458,5,31,
        0,0,458,460,1,0,0,0,459,456,1,0,0,0,460,463,1,0,0,0,461,459,1,0,
        0,0,461,462,1,0,0,0,462,464,1,0,0,0,463,461,1,0,0,0,464,466,3,54,
        27,0,465,461,1,0,0,0,465,466,1,0,0,0,466,57,1,0,0,0,467,468,6,29,
        -1,0,468,475,3,42,21,0,469,470,3,42,21,0,470,471,5,35,0,0,471,472,
        3,56,28,0,472,473,5,34,0,0,473,475,1,0,0,0,474,467,1,0,0,0,474,469,
        1,0,0,0,475,483,1,0,0,0,476,477,10,1,0,0,477,478,5,25,0,0,478,479,
        3,40,20,0,479,480,5,26,0,0,480,482,1,0,0,0,481,476,1,0,0,0,482,485,
        1,0,0,0,483,481,1,0,0,0,483,484,1,0,0,0,484,59,1,0,0,0,485,483,1,
        0,0,0,33,63,67,73,76,80,97,101,106,130,137,141,150,156,166,175,179,
        183,210,214,227,243,250,356,420,422,429,435,450,454,461,465,474,
        483
    ]

class RiddleParser ( Parser ):

    grammarFileName = "RiddleParser.g4"

    atn = ATNDeserializer().deserialize(serializedATN())

    decisionsToDFA = [ DFA(ds, i) for i, ds in enumerate(atn.decisionToState) ]

    sharedContextCache = PredictionContextCache()

    literalNames = [ "<INVALID>", "'var'", "'val'", "'for'", "'while'", 
                     "'if'", "'else'", "'fun'", "'return'", "'import'", 
                     "'package'", "'class'", "'public'", "'protected'", 
                     "'Private'", "'override'", "'true'", "'false'", "'static'", 
                     "'const'", "'null'", "'try'", "'catch'", "'('", "')'", 
                     "'['", "']'", "'{'", "'}'", "':'", "';'", "','", "'=='", 
                     "'='", "'>'", "'<'", "'<<'", "'>>'", "'>>>'", "'+'", 
                     "'-'", "'*'", "'/'", "'%'", "'!'", "'&'", "'|'", "'^'", 
                     "'.'", "'\"'", "'''", "'\\n'" ]

    symbolicNames = [ "<INVALID>", "Var", "Val", "For", "While", "If", "Else", 
                      "Func", "Return", "Import", "Package", "Class", "Public", 
                      "Protected", "Private", "Override", "True_", "False_", 
                      "Static", "Const", "Null", "Try", "Catch", "LeftBracket", 
                      "RightBracket", "LeftSquare", "RightSquare", "LeftCurly", 
                      "RightCurly", "Colon", "Semi", "Comma", "Equal", "Assign", 
                      "Greater", "Less", "LeftLeft", "RightRight", "RightRightRight", 
                      "Add", "Sub", "Star", "Div", "Mod", "Not", "And", 
                      "Or", "Xor", "Dot", "DoubleQuotes", "Quotes", "Endl", 
                      "Identifier", "Hexadecimal", "Decimal", "Octal", "Binary", 
                      "Float", "IntegerSequence", "HEX_DIGIT", "OCTAL_DIGIT", 
                      "BINARY_DIGIT", "DIGIT", "STRING", "LINE_COMMENT", 
                      "BLOCK_COMMENT", "WHITESPACE" ]

    RULE_program = 0
    RULE_null_cnt = 1
    RULE_statement_ed = 2
    RULE_statement = 3
    RULE_bodyExpr = 4
    RULE_packStatement = 5
    RULE_importStatement = 6
    RULE_varDefineStatement = 7
    RULE_argsExpr = 8
    RULE_defineArgs = 9
    RULE_funcDefine = 10
    RULE_forStatement = 11
    RULE_whileStatement = 12
    RULE_ifStatement = 13
    RULE_returnStatement = 14
    RULE_classDefine = 15
    RULE_tryExpr = 16
    RULE_catchExpr = 17
    RULE_exprPtr = 18
    RULE_exprPtrParser = 19
    RULE_expression = 20
    RULE_id = 21
    RULE_number = 22
    RULE_boolean = 23
    RULE_string = 24
    RULE_float = 25
    RULE_integer = 26
    RULE_templateArg = 27
    RULE_templateArgs = 28
    RULE_typeName = 29

    ruleNames =  [ "program", "null_cnt", "statement_ed", "statement", "bodyExpr", 
                   "packStatement", "importStatement", "varDefineStatement", 
                   "argsExpr", "defineArgs", "funcDefine", "forStatement", 
                   "whileStatement", "ifStatement", "returnStatement", "classDefine", 
                   "tryExpr", "catchExpr", "exprPtr", "exprPtrParser", "expression", 
                   "id", "number", "boolean", "string", "float", "integer", 
                   "templateArg", "templateArgs", "typeName" ]

    EOF = Token.EOF
    Var=1
    Val=2
    For=3
    While=4
    If=5
    Else=6
    Func=7
    Return=8
    Import=9
    Package=10
    Class=11
    Public=12
    Protected=13
    Private=14
    Override=15
    True_=16
    False_=17
    Static=18
    Const=19
    Null=20
    Try=21
    Catch=22
    LeftBracket=23
    RightBracket=24
    LeftSquare=25
    RightSquare=26
    LeftCurly=27
    RightCurly=28
    Colon=29
    Semi=30
    Comma=31
    Equal=32
    Assign=33
    Greater=34
    Less=35
    LeftLeft=36
    RightRight=37
    RightRightRight=38
    Add=39
    Sub=40
    Star=41
    Div=42
    Mod=43
    Not=44
    And=45
    Or=46
    Xor=47
    Dot=48
    DoubleQuotes=49
    Quotes=50
    Endl=51
    Identifier=52
    Hexadecimal=53
    Decimal=54
    Octal=55
    Binary=56
    Float=57
    IntegerSequence=58
    HEX_DIGIT=59
    OCTAL_DIGIT=60
    BINARY_DIGIT=61
    DIGIT=62
    STRING=63
    LINE_COMMENT=64
    BLOCK_COMMENT=65
    WHITESPACE=66

    def __init__(self, input:TokenStream, output:TextIO = sys.stdout):
        super().__init__(input, output)
        self.checkVersion("4.13.2")
        self._interp = ParserATNSimulator(self, self.atn, self.decisionsToDFA, self.sharedContextCache)
        self._predicates = None




    class ProgramContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def statement_ed(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.Statement_edContext)
            else:
                return self.getTypedRuleContext(RiddleParser.Statement_edContext,i)


        def EOF(self):
            return self.getToken(RiddleParser.EOF, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_program

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterProgram" ):
                listener.enterProgram(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitProgram" ):
                listener.exitProgram(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitProgram" ):
                return visitor.visitProgram(self)
            else:
                return visitor.visitChildren(self)




    def program(self):

        localctx = RiddleParser.ProgramContext(self, self._ctx, self.state)
        self.enterRule(localctx, 0, self.RULE_program)
        self._la = 0 # Token type
        try:
            self.state = 67
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,1,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 63
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & -8937374183483830342) != 0):
                    self.state = 60
                    self.statement_ed()
                    self.state = 65
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 66
                self.match(RiddleParser.EOF)
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Null_cntContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def Semi(self):
            return self.getToken(RiddleParser.Semi, 0)

        def Endl(self):
            return self.getToken(RiddleParser.Endl, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_null_cnt

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNull_cnt" ):
                listener.enterNull_cnt(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNull_cnt" ):
                listener.exitNull_cnt(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNull_cnt" ):
                return visitor.visitNull_cnt(self)
            else:
                return visitor.visitChildren(self)




    def null_cnt(self):

        localctx = RiddleParser.Null_cntContext(self, self._ctx, self.state)
        self.enterRule(localctx, 2, self.RULE_null_cnt)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 69
            _la = self._input.LA(1)
            if not(_la==30 or _la==51):
                self._errHandler.recoverInline(self)
            else:
                self._errHandler.reportMatch(self)
                self.consume()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class Statement_edContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def statement(self):
            return self.getTypedRuleContext(RiddleParser.StatementContext,0)


        def Semi(self):
            return self.getToken(RiddleParser.Semi, 0)

        def Endl(self):
            return self.getToken(RiddleParser.Endl, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_statement_ed

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStatement_ed" ):
                listener.enterStatement_ed(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStatement_ed" ):
                listener.exitStatement_ed(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitStatement_ed" ):
                return visitor.visitStatement_ed(self)
            else:
                return visitor.visitChildren(self)




    def statement_ed(self):

        localctx = RiddleParser.Statement_edContext(self, self._ctx, self.state)
        self.enterRule(localctx, 4, self.RULE_statement_ed)
        try:
            self.state = 80
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [1, 3, 4, 5, 7, 8, 9, 10, 11, 16, 17, 20, 21, 23, 27, 35, 39, 40, 44, 52, 53, 54, 55, 56, 57, 63]:
                self.enterOuterAlt(localctx, 1)
                self.state = 71
                self.statement()
                self.state = 73
                self._errHandler.sync(self)
                la_ = self._interp.adaptivePredict(self._input,2,self._ctx)
                if la_ == 1:
                    self.state = 72
                    self.match(RiddleParser.Semi)


                self.state = 76
                self._errHandler.sync(self)
                la_ = self._interp.adaptivePredict(self._input,3,self._ctx)
                if la_ == 1:
                    self.state = 75
                    self.match(RiddleParser.Endl)


                pass
            elif token in [30]:
                self.enterOuterAlt(localctx, 2)
                self.state = 78
                self.match(RiddleParser.Semi)
                pass
            elif token in [51]:
                self.enterOuterAlt(localctx, 3)
                self.state = 79
                self.match(RiddleParser.Endl)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def packStatement(self):
            return self.getTypedRuleContext(RiddleParser.PackStatementContext,0)


        def importStatement(self):
            return self.getTypedRuleContext(RiddleParser.ImportStatementContext,0)


        def classDefine(self):
            return self.getTypedRuleContext(RiddleParser.ClassDefineContext,0)


        def funcDefine(self):
            return self.getTypedRuleContext(RiddleParser.FuncDefineContext,0)


        def varDefineStatement(self):
            return self.getTypedRuleContext(RiddleParser.VarDefineStatementContext,0)


        def forStatement(self):
            return self.getTypedRuleContext(RiddleParser.ForStatementContext,0)


        def whileStatement(self):
            return self.getTypedRuleContext(RiddleParser.WhileStatementContext,0)


        def ifStatement(self):
            return self.getTypedRuleContext(RiddleParser.IfStatementContext,0)


        def returnStatement(self):
            return self.getTypedRuleContext(RiddleParser.ReturnStatementContext,0)


        def tryExpr(self):
            return self.getTypedRuleContext(RiddleParser.TryExprContext,0)


        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def LeftCurly(self):
            return self.getToken(RiddleParser.LeftCurly, 0)

        def RightCurly(self):
            return self.getToken(RiddleParser.RightCurly, 0)

        def statement_ed(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.Statement_edContext)
            else:
                return self.getTypedRuleContext(RiddleParser.Statement_edContext,i)


        def getRuleIndex(self):
            return RiddleParser.RULE_statement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStatement" ):
                listener.enterStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStatement" ):
                listener.exitStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitStatement" ):
                return visitor.visitStatement(self)
            else:
                return visitor.visitChildren(self)




    def statement(self):

        localctx = RiddleParser.StatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 6, self.RULE_statement)
        self._la = 0 # Token type
        try:
            self.state = 101
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [10]:
                self.enterOuterAlt(localctx, 1)
                self.state = 82
                self.packStatement()
                pass
            elif token in [9]:
                self.enterOuterAlt(localctx, 2)
                self.state = 83
                self.importStatement()
                pass
            elif token in [11]:
                self.enterOuterAlt(localctx, 3)
                self.state = 84
                self.classDefine()
                pass
            elif token in [7]:
                self.enterOuterAlt(localctx, 4)
                self.state = 85
                self.funcDefine()
                pass
            elif token in [1]:
                self.enterOuterAlt(localctx, 5)
                self.state = 86
                self.varDefineStatement()
                pass
            elif token in [3]:
                self.enterOuterAlt(localctx, 6)
                self.state = 87
                self.forStatement()
                pass
            elif token in [4]:
                self.enterOuterAlt(localctx, 7)
                self.state = 88
                self.whileStatement()
                pass
            elif token in [5]:
                self.enterOuterAlt(localctx, 8)
                self.state = 89
                self.ifStatement()
                pass
            elif token in [8]:
                self.enterOuterAlt(localctx, 9)
                self.state = 90
                self.returnStatement()
                pass
            elif token in [21]:
                self.enterOuterAlt(localctx, 10)
                self.state = 91
                self.tryExpr()
                pass
            elif token in [16, 17, 20, 23, 35, 39, 40, 44, 52, 53, 54, 55, 56, 57, 63]:
                self.enterOuterAlt(localctx, 11)
                self.state = 92
                self.expression(0)
                pass
            elif token in [27]:
                self.enterOuterAlt(localctx, 12)
                self.state = 93
                self.match(RiddleParser.LeftCurly)
                self.state = 97
                self._errHandler.sync(self)
                _la = self._input.LA(1)
                while (((_la) & ~0x3f) == 0 and ((1 << _la) & -8937374183483830342) != 0):
                    self.state = 94
                    self.statement_ed()
                    self.state = 99
                    self._errHandler.sync(self)
                    _la = self._input.LA(1)

                self.state = 100
                self.match(RiddleParser.RightCurly)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class BodyExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def statement_ed(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.Statement_edContext)
            else:
                return self.getTypedRuleContext(RiddleParser.Statement_edContext,i)


        def getRuleIndex(self):
            return RiddleParser.RULE_bodyExpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBodyExpr" ):
                listener.enterBodyExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBodyExpr" ):
                listener.exitBodyExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBodyExpr" ):
                return visitor.visitBodyExpr(self)
            else:
                return visitor.visitChildren(self)




    def bodyExpr(self):

        localctx = RiddleParser.BodyExprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 8, self.RULE_bodyExpr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 106
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            while (((_la) & ~0x3f) == 0 and ((1 << _la) & -8937374183483830342) != 0):
                self.state = 103
                self.statement_ed()
                self.state = 108
                self._errHandler.sync(self)
                _la = self._input.LA(1)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class PackStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.packName = None # IdContext

        def Package(self):
            return self.getToken(RiddleParser.Package, 0)

        def id_(self):
            return self.getTypedRuleContext(RiddleParser.IdContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_packStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPackStatement" ):
                listener.enterPackStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPackStatement" ):
                listener.exitPackStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitPackStatement" ):
                return visitor.visitPackStatement(self)
            else:
                return visitor.visitChildren(self)




    def packStatement(self):

        localctx = RiddleParser.PackStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 10, self.RULE_packStatement)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 109
            self.match(RiddleParser.Package)
            self.state = 110
            localctx.packName = self.id_()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ImportStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.libName = None # IdContext

        def Import(self):
            return self.getToken(RiddleParser.Import, 0)

        def id_(self):
            return self.getTypedRuleContext(RiddleParser.IdContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_importStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterImportStatement" ):
                listener.enterImportStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitImportStatement" ):
                listener.exitImportStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitImportStatement" ):
                return visitor.visitImportStatement(self)
            else:
                return visitor.visitChildren(self)




    def importStatement(self):

        localctx = RiddleParser.ImportStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 12, self.RULE_importStatement)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 112
            self.match(RiddleParser.Import)
            self.state = 113
            localctx.libName = self.id_()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class VarDefineStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.name = None # Token
            self.type_ = None # TypeNameContext
            self.value = None # ExpressionContext

        def Var(self):
            return self.getToken(RiddleParser.Var, 0)

        def Colon(self):
            return self.getToken(RiddleParser.Colon, 0)

        def Identifier(self):
            return self.getToken(RiddleParser.Identifier, 0)

        def typeName(self):
            return self.getTypedRuleContext(RiddleParser.TypeNameContext,0)


        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_varDefineStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterVarDefineStatement" ):
                listener.enterVarDefineStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitVarDefineStatement" ):
                listener.exitVarDefineStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitVarDefineStatement" ):
                return visitor.visitVarDefineStatement(self)
            else:
                return visitor.visitChildren(self)




    def varDefineStatement(self):

        localctx = RiddleParser.VarDefineStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 14, self.RULE_varDefineStatement)
        try:
            self.state = 130
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,8,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 115
                self.match(RiddleParser.Var)
                self.state = 116
                localctx.name = self.match(RiddleParser.Identifier)
                self.state = 117
                self.match(RiddleParser.Colon)
                self.state = 118
                localctx.type_ = self.typeName(0)
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 119
                self.match(RiddleParser.Var)
                self.state = 120
                localctx.name = self.match(RiddleParser.Identifier)
                self.state = 121
                self.match(RiddleParser.Assign)
                self.state = 122
                localctx.value = self.expression(0)
                pass

            elif la_ == 3:
                self.enterOuterAlt(localctx, 3)
                self.state = 123
                self.match(RiddleParser.Var)
                self.state = 124
                localctx.name = self.match(RiddleParser.Identifier)
                self.state = 125
                self.match(RiddleParser.Colon)
                self.state = 126
                localctx.type_ = self.typeName(0)
                self.state = 127
                self.match(RiddleParser.Assign)
                self.state = 128
                localctx.value = self.expression(0)
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ArgsExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def Comma(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Comma)
            else:
                return self.getToken(RiddleParser.Comma, i)

        def getRuleIndex(self):
            return RiddleParser.RULE_argsExpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterArgsExpr" ):
                listener.enterArgsExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitArgsExpr" ):
                listener.exitArgsExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitArgsExpr" ):
                return visitor.visitArgsExpr(self)
            else:
                return visitor.visitChildren(self)




    def argsExpr(self):

        localctx = RiddleParser.ArgsExprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 16, self.RULE_argsExpr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 141
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & -8939625984507576320) != 0):
                self.state = 137
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,9,self._ctx)
                while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                    if _alt==1:
                        self.state = 132
                        self.expression(0)
                        self.state = 133
                        self.match(RiddleParser.Comma) 
                    self.state = 139
                    self._errHandler.sync(self)
                    _alt = self._interp.adaptivePredict(self._input,9,self._ctx)

                self.state = 140
                self.expression(0)


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class DefineArgsContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def Identifier(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Identifier)
            else:
                return self.getToken(RiddleParser.Identifier, i)

        def Colon(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Colon)
            else:
                return self.getToken(RiddleParser.Colon, i)

        def typeName(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.TypeNameContext)
            else:
                return self.getTypedRuleContext(RiddleParser.TypeNameContext,i)


        def Comma(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Comma)
            else:
                return self.getToken(RiddleParser.Comma, i)

        def getRuleIndex(self):
            return RiddleParser.RULE_defineArgs

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDefineArgs" ):
                listener.enterDefineArgs(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDefineArgs" ):
                listener.exitDefineArgs(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitDefineArgs" ):
                return visitor.visitDefineArgs(self)
            else:
                return visitor.visitChildren(self)




    def defineArgs(self):

        localctx = RiddleParser.DefineArgsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 18, self.RULE_defineArgs)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 156
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if _la==52:
                self.state = 150
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,11,self._ctx)
                while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                    if _alt==1:
                        self.state = 143
                        self.match(RiddleParser.Identifier)
                        self.state = 144
                        self.match(RiddleParser.Colon)
                        self.state = 145
                        self.typeName(0)
                        self.state = 146
                        self.match(RiddleParser.Comma) 
                    self.state = 152
                    self._errHandler.sync(self)
                    _alt = self._interp.adaptivePredict(self._input,11,self._ctx)

                self.state = 153
                self.match(RiddleParser.Identifier)
                self.state = 154
                self.match(RiddleParser.Colon)
                self.state = 155
                self.typeName(0)


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class FuncDefineContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.funcName = None # Token
            self.args = None # DefineArgsContext
            self.returnType = None # TypeNameContext
            self.body = None # BodyExprContext

        def Func(self):
            return self.getToken(RiddleParser.Func, 0)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)

        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)

        def LeftCurly(self):
            return self.getToken(RiddleParser.LeftCurly, 0)

        def RightCurly(self):
            return self.getToken(RiddleParser.RightCurly, 0)

        def Identifier(self):
            return self.getToken(RiddleParser.Identifier, 0)

        def defineArgs(self):
            return self.getTypedRuleContext(RiddleParser.DefineArgsContext,0)


        def bodyExpr(self):
            return self.getTypedRuleContext(RiddleParser.BodyExprContext,0)


        def Sub(self):
            return self.getToken(RiddleParser.Sub, 0)

        def Greater(self):
            return self.getToken(RiddleParser.Greater, 0)

        def typeName(self):
            return self.getTypedRuleContext(RiddleParser.TypeNameContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_funcDefine

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFuncDefine" ):
                listener.enterFuncDefine(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFuncDefine" ):
                listener.exitFuncDefine(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitFuncDefine" ):
                return visitor.visitFuncDefine(self)
            else:
                return visitor.visitChildren(self)




    def funcDefine(self):

        localctx = RiddleParser.FuncDefineContext(self, self._ctx, self.state)
        self.enterRule(localctx, 20, self.RULE_funcDefine)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 158
            self.match(RiddleParser.Func)
            self.state = 159
            localctx.funcName = self.match(RiddleParser.Identifier)
            self.state = 160
            self.match(RiddleParser.LeftBracket)
            self.state = 161
            localctx.args = self.defineArgs()
            self.state = 162
            self.match(RiddleParser.RightBracket)
            self.state = 166
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if _la==40:
                self.state = 163
                self.match(RiddleParser.Sub)
                self.state = 164
                self.match(RiddleParser.Greater)
                self.state = 165
                localctx.returnType = self.typeName(0)


            self.state = 168
            self.match(RiddleParser.LeftCurly)
            self.state = 169
            localctx.body = self.bodyExpr()
            self.state = 170
            self.match(RiddleParser.RightCurly)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ForStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.init = None # VarDefineStatementContext
            self.termCond = None # ExpressionContext
            self.selfVar = None # StatementContext
            self.body = None # Statement_edContext

        def For(self):
            return self.getToken(RiddleParser.For, 0)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)

        def Semi(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Semi)
            else:
                return self.getToken(RiddleParser.Semi, i)

        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)

        def statement_ed(self):
            return self.getTypedRuleContext(RiddleParser.Statement_edContext,0)


        def varDefineStatement(self):
            return self.getTypedRuleContext(RiddleParser.VarDefineStatementContext,0)


        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def statement(self):
            return self.getTypedRuleContext(RiddleParser.StatementContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_forStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterForStatement" ):
                listener.enterForStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitForStatement" ):
                listener.exitForStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitForStatement" ):
                return visitor.visitForStatement(self)
            else:
                return visitor.visitChildren(self)




    def forStatement(self):

        localctx = RiddleParser.ForStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 22, self.RULE_forStatement)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 172
            self.match(RiddleParser.For)
            self.state = 173
            self.match(RiddleParser.LeftBracket)
            self.state = 175
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if _la==1:
                self.state = 174
                localctx.init = self.varDefineStatement()


            self.state = 177
            self.match(RiddleParser.Semi)
            self.state = 179
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & -8939625984507576320) != 0):
                self.state = 178
                localctx.termCond = self.expression(0)


            self.state = 181
            self.match(RiddleParser.Semi)
            self.state = 183
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & -8939625984371257414) != 0):
                self.state = 182
                localctx.selfVar = self.statement()


            self.state = 185
            self.match(RiddleParser.RightBracket)
            self.state = 186
            localctx.body = self.statement_ed()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class WhileStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.runCond = None # ExpressionContext
            self.body = None # Statement_edContext

        def While(self):
            return self.getToken(RiddleParser.While, 0)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)

        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def statement_ed(self):
            return self.getTypedRuleContext(RiddleParser.Statement_edContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_whileStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterWhileStatement" ):
                listener.enterWhileStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitWhileStatement" ):
                listener.exitWhileStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitWhileStatement" ):
                return visitor.visitWhileStatement(self)
            else:
                return visitor.visitChildren(self)




    def whileStatement(self):

        localctx = RiddleParser.WhileStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 24, self.RULE_whileStatement)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 188
            self.match(RiddleParser.While)
            self.state = 189
            self.match(RiddleParser.LeftBracket)
            self.state = 190
            localctx.runCond = self.expression(0)
            self.state = 191
            self.match(RiddleParser.RightBracket)
            self.state = 192
            localctx.body = self.statement_ed()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class IfStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.hasElse = None
            self.cond = None # ExpressionContext
            self.body = None # Statement_edContext
            self.elseBody = None # Statement_edContext

        def If(self):
            return self.getToken(RiddleParser.If, 0)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)

        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def statement_ed(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.Statement_edContext)
            else:
                return self.getTypedRuleContext(RiddleParser.Statement_edContext,i)


        def Else(self):
            return self.getToken(RiddleParser.Else, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_ifStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterIfStatement" ):
                listener.enterIfStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitIfStatement" ):
                listener.exitIfStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitIfStatement" ):
                return visitor.visitIfStatement(self)
            else:
                return visitor.visitChildren(self)




    def ifStatement(self):

        localctx = RiddleParser.IfStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 26, self.RULE_ifStatement)
        try:
            self.state = 210
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,17,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 194
                self.match(RiddleParser.If)
                self.state = 195
                self.match(RiddleParser.LeftBracket)
                self.state = 196
                localctx.cond = self.expression(0)
                self.state = 197
                self.match(RiddleParser.RightBracket)
                self.state = 198
                localctx.body = self.statement_ed()
                localctx.hasElse=False
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 201
                self.match(RiddleParser.If)
                self.state = 202
                self.match(RiddleParser.LeftBracket)
                self.state = 203
                localctx.cond = self.expression(0)
                self.state = 204
                self.match(RiddleParser.RightBracket)
                self.state = 205
                localctx.body = self.statement_ed()
                self.state = 206
                self.match(RiddleParser.Else)
                self.state = 207
                localctx.elseBody = self.statement_ed()
                localctx.hasElse=True
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ReturnStatementContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.result = None # Statement_edContext

        def Return(self):
            return self.getToken(RiddleParser.Return, 0)

        def statement_ed(self):
            return self.getTypedRuleContext(RiddleParser.Statement_edContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_returnStatement

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterReturnStatement" ):
                listener.enterReturnStatement(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitReturnStatement" ):
                listener.exitReturnStatement(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitReturnStatement" ):
                return visitor.visitReturnStatement(self)
            else:
                return visitor.visitChildren(self)




    def returnStatement(self):

        localctx = RiddleParser.ReturnStatementContext(self, self._ctx, self.state)
        self.enterRule(localctx, 28, self.RULE_returnStatement)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 212
            self.match(RiddleParser.Return)
            self.state = 214
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,18,self._ctx)
            if la_ == 1:
                self.state = 213
                localctx.result = self.statement_ed()


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ClassDefineContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.className = None # IdContext
            self.body = None # BodyExprContext

        def Class(self):
            return self.getToken(RiddleParser.Class, 0)

        def LeftCurly(self):
            return self.getToken(RiddleParser.LeftCurly, 0)

        def RightCurly(self):
            return self.getToken(RiddleParser.RightCurly, 0)

        def id_(self):
            return self.getTypedRuleContext(RiddleParser.IdContext,0)


        def bodyExpr(self):
            return self.getTypedRuleContext(RiddleParser.BodyExprContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_classDefine

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterClassDefine" ):
                listener.enterClassDefine(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitClassDefine" ):
                listener.exitClassDefine(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitClassDefine" ):
                return visitor.visitClassDefine(self)
            else:
                return visitor.visitChildren(self)




    def classDefine(self):

        localctx = RiddleParser.ClassDefineContext(self, self._ctx, self.state)
        self.enterRule(localctx, 30, self.RULE_classDefine)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 216
            self.match(RiddleParser.Class)
            self.state = 217
            localctx.className = self.id_()
            self.state = 218
            self.match(RiddleParser.LeftCurly)
            self.state = 219
            localctx.body = self.bodyExpr()
            self.state = 220
            self.match(RiddleParser.RightCurly)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TryExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.tryBody = None # BodyExprContext

        def Try(self):
            return self.getToken(RiddleParser.Try, 0)

        def LeftCurly(self):
            return self.getToken(RiddleParser.LeftCurly, 0)

        def RightCurly(self):
            return self.getToken(RiddleParser.RightCurly, 0)

        def catchExpr(self):
            return self.getTypedRuleContext(RiddleParser.CatchExprContext,0)


        def bodyExpr(self):
            return self.getTypedRuleContext(RiddleParser.BodyExprContext,0)


        def null_cnt(self):
            return self.getTypedRuleContext(RiddleParser.Null_cntContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_tryExpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTryExpr" ):
                listener.enterTryExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTryExpr" ):
                listener.exitTryExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitTryExpr" ):
                return visitor.visitTryExpr(self)
            else:
                return visitor.visitChildren(self)




    def tryExpr(self):

        localctx = RiddleParser.TryExprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 32, self.RULE_tryExpr)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 222
            self.match(RiddleParser.Try)
            self.state = 223
            self.match(RiddleParser.LeftCurly)
            self.state = 224
            localctx.tryBody = self.bodyExpr()
            self.state = 225
            self.match(RiddleParser.RightCurly)
            self.state = 227
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if _la==30 or _la==51:
                self.state = 226
                self.null_cnt()


            self.state = 229
            self.catchExpr()
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class CatchExprContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def Catch(self):
            return self.getToken(RiddleParser.Catch, 0)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)

        def varDefineStatement(self):
            return self.getTypedRuleContext(RiddleParser.VarDefineStatementContext,0)


        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_catchExpr

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCatchExpr" ):
                listener.enterCatchExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCatchExpr" ):
                listener.exitCatchExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCatchExpr" ):
                return visitor.visitCatchExpr(self)
            else:
                return visitor.visitChildren(self)




    def catchExpr(self):

        localctx = RiddleParser.CatchExprContext(self, self._ctx, self.state)
        self.enterRule(localctx, 34, self.RULE_catchExpr)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 231
            self.match(RiddleParser.Catch)
            self.state = 232
            self.match(RiddleParser.LeftBracket)
            self.state = 233
            self.varDefineStatement()
            self.state = 234
            self.match(RiddleParser.RightBracket)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExprPtrContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return RiddleParser.RULE_exprPtr

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)


    class FuncExprContext(ExprPtrContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExprPtrContext
            super().__init__(parser)
            self.funcName = None # Token
            self.args = None # ArgsExprContext
            self.copyFrom(ctx)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)
        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)
        def Identifier(self):
            return self.getToken(RiddleParser.Identifier, 0)
        def argsExpr(self):
            return self.getTypedRuleContext(RiddleParser.ArgsExprContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFuncExpr" ):
                listener.enterFuncExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFuncExpr" ):
                listener.exitFuncExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitFuncExpr" ):
                return visitor.visitFuncExpr(self)
            else:
                return visitor.visitChildren(self)


    class ObjectExprContext(ExprPtrContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExprPtrContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def Identifier(self):
            return self.getToken(RiddleParser.Identifier, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterObjectExpr" ):
                listener.enterObjectExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitObjectExpr" ):
                listener.exitObjectExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitObjectExpr" ):
                return visitor.visitObjectExpr(self)
            else:
                return visitor.visitChildren(self)


    class BlendExprContext(ExprPtrContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExprPtrContext
            super().__init__(parser)
            self.parent = None # ExprPtrContext
            self.child = None # ExprPtrContext
            self.copyFrom(ctx)

        def Dot(self):
            return self.getToken(RiddleParser.Dot, 0)
        def exprPtr(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExprPtrContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExprPtrContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBlendExpr" ):
                listener.enterBlendExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBlendExpr" ):
                listener.exitBlendExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBlendExpr" ):
                return visitor.visitBlendExpr(self)
            else:
                return visitor.visitChildren(self)



    def exprPtr(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = RiddleParser.ExprPtrContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 36
        self.enterRecursionRule(localctx, 36, self.RULE_exprPtr, _p)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 243
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,20,self._ctx)
            if la_ == 1:
                localctx = RiddleParser.FuncExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 237
                localctx.funcName = self.match(RiddleParser.Identifier)
                self.state = 238
                self.match(RiddleParser.LeftBracket)
                self.state = 239
                localctx.args = self.argsExpr()
                self.state = 240
                self.match(RiddleParser.RightBracket)
                pass

            elif la_ == 2:
                localctx = RiddleParser.ObjectExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 242
                self.match(RiddleParser.Identifier)
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 250
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,21,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    localctx = RiddleParser.BlendExprContext(self, RiddleParser.ExprPtrContext(self, _parentctx, _parentState))
                    localctx.parent = _prevctx
                    self.pushNewRecursionContext(localctx, _startState, self.RULE_exprPtr)
                    self.state = 245
                    if not self.precpred(self._ctx, 1):
                        from antlr4.error.Errors import FailedPredicateException
                        raise FailedPredicateException(self, "self.precpred(self._ctx, 1)")
                    self.state = 246
                    self.match(RiddleParser.Dot)
                    self.state = 247
                    localctx.child = self.exprPtr(2) 
                self.state = 252
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,21,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class ExprPtrParserContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def exprPtr(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_exprPtrParser

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterExprPtrParser" ):
                listener.enterExprPtrParser(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitExprPtrParser" ):
                listener.exitExprPtrParser(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitExprPtrParser" ):
                return visitor.visitExprPtrParser(self)
            else:
                return visitor.visitChildren(self)




    def exprPtrParser(self):

        localctx = RiddleParser.ExprPtrParserContext(self, self._ctx, self.state)
        self.enterRule(localctx, 38, self.RULE_exprPtrParser)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 253
            self.exprPtr(0)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class ExpressionContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser


        def getRuleIndex(self):
            return RiddleParser.RULE_expression

     
        def copyFrom(self, ctx:ParserRuleContext):
            super().copyFrom(ctx)


    class SelfSubRightExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExprPtrParserContext
            self.copyFrom(ctx)

        def Sub(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Sub)
            else:
                return self.getToken(RiddleParser.Sub, i)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSelfSubRightExpr" ):
                listener.enterSelfSubRightExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSelfSubRightExpr" ):
                listener.exitSelfSubRightExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSelfSubRightExpr" ):
                return visitor.visitSelfSubRightExpr(self)
            else:
                return visitor.visitChildren(self)


    class AndAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def And(self):
            return self.getToken(RiddleParser.And, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAndAssignExpr" ):
                listener.enterAndAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAndAssignExpr" ):
                listener.exitAndAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAndAssignExpr" ):
                return visitor.visitAndAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class NullExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def Null(self):
            return self.getToken(RiddleParser.Null, 0)

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNullExpr" ):
                listener.enterNullExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNullExpr" ):
                listener.exitNullExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNullExpr" ):
                return visitor.visitNullExpr(self)
            else:
                return visitor.visitChildren(self)


    class ModExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Mod(self):
            return self.getToken(RiddleParser.Mod, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterModExpr" ):
                listener.enterModExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitModExpr" ):
                listener.exitModExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitModExpr" ):
                return visitor.visitModExpr(self)
            else:
                return visitor.visitChildren(self)


    class CastExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.type_ = None # TypeNameContext
            self.value = None # ExprPtrParserContext
            self.copyFrom(ctx)

        def Less(self):
            return self.getToken(RiddleParser.Less, 0)
        def Greater(self):
            return self.getToken(RiddleParser.Greater, 0)
        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)
        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)
        def typeName(self):
            return self.getTypedRuleContext(RiddleParser.TypeNameContext,0)

        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterCastExpr" ):
                listener.enterCastExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitCastExpr" ):
                listener.exitCastExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitCastExpr" ):
                return visitor.visitCastExpr(self)
            else:
                return visitor.visitChildren(self)


    class LShrExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def RightRightRight(self):
            return self.getToken(RiddleParser.RightRightRight, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLShrExpr" ):
                listener.enterLShrExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLShrExpr" ):
                listener.exitLShrExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLShrExpr" ):
                return visitor.visitLShrExpr(self)
            else:
                return visitor.visitChildren(self)


    class NegativeExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExpressionContext
            self.copyFrom(ctx)

        def Sub(self):
            return self.getToken(RiddleParser.Sub, 0)
        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNegativeExpr" ):
                listener.enterNegativeExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNegativeExpr" ):
                listener.exitNegativeExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNegativeExpr" ):
                return visitor.visitNegativeExpr(self)
            else:
                return visitor.visitChildren(self)


    class NumberExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def number(self):
            return self.getTypedRuleContext(RiddleParser.NumberContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNumberExpr" ):
                listener.enterNumberExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNumberExpr" ):
                listener.exitNumberExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNumberExpr" ):
                return visitor.visitNumberExpr(self)
            else:
                return visitor.visitChildren(self)


    class MulAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Star(self):
            return self.getToken(RiddleParser.Star, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMulAssignExpr" ):
                listener.enterMulAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMulAssignExpr" ):
                listener.exitMulAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitMulAssignExpr" ):
                return visitor.visitMulAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class XorAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Xor(self):
            return self.getToken(RiddleParser.Xor, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterXorAssignExpr" ):
                listener.enterXorAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitXorAssignExpr" ):
                listener.exitXorAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitXorAssignExpr" ):
                return visitor.visitXorAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class SubExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Sub(self):
            return self.getToken(RiddleParser.Sub, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSubExpr" ):
                listener.enterSubExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSubExpr" ):
                listener.exitSubExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSubExpr" ):
                return visitor.visitSubExpr(self)
            else:
                return visitor.visitChildren(self)


    class GreaterEqualExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Greater(self):
            return self.getToken(RiddleParser.Greater, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterGreaterEqualExpr" ):
                listener.enterGreaterEqualExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitGreaterEqualExpr" ):
                listener.exitGreaterEqualExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitGreaterEqualExpr" ):
                return visitor.visitGreaterEqualExpr(self)
            else:
                return visitor.visitChildren(self)


    class AddAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Add(self):
            return self.getToken(RiddleParser.Add, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAddAssignExpr" ):
                listener.enterAddAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAddAssignExpr" ):
                listener.exitAddAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAddAssignExpr" ):
                return visitor.visitAddAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class OrAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Or(self):
            return self.getToken(RiddleParser.Or, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterOrAssignExpr" ):
                listener.enterOrAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitOrAssignExpr" ):
                listener.exitOrAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitOrAssignExpr" ):
                return visitor.visitOrAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class BitXorExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Xor(self):
            return self.getToken(RiddleParser.Xor, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBitXorExpr" ):
                listener.enterBitXorExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBitXorExpr" ):
                listener.exitBitXorExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBitXorExpr" ):
                return visitor.visitBitXorExpr(self)
            else:
                return visitor.visitChildren(self)


    class PtrExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def exprPtr(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPtrExpr" ):
                listener.enterPtrExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPtrExpr" ):
                listener.exitPtrExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitPtrExpr" ):
                return visitor.visitPtrExpr(self)
            else:
                return visitor.visitChildren(self)


    class StringExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def string(self):
            return self.getTypedRuleContext(RiddleParser.StringContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterStringExpr" ):
                listener.enterStringExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitStringExpr" ):
                listener.exitStringExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitStringExpr" ):
                return visitor.visitStringExpr(self)
            else:
                return visitor.visitChildren(self)


    class GreaterExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Greater(self):
            return self.getToken(RiddleParser.Greater, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterGreaterExpr" ):
                listener.enterGreaterExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitGreaterExpr" ):
                listener.exitGreaterExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitGreaterExpr" ):
                return visitor.visitGreaterExpr(self)
            else:
                return visitor.visitChildren(self)


    class DivAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Div(self):
            return self.getToken(RiddleParser.Div, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDivAssignExpr" ):
                listener.enterDivAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDivAssignExpr" ):
                listener.exitDivAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitDivAssignExpr" ):
                return visitor.visitDivAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class MulExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Star(self):
            return self.getToken(RiddleParser.Star, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterMulExpr" ):
                listener.enterMulExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitMulExpr" ):
                listener.exitMulExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitMulExpr" ):
                return visitor.visitMulExpr(self)
            else:
                return visitor.visitChildren(self)


    class NotEqualExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Not(self):
            return self.getToken(RiddleParser.Not, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNotEqualExpr" ):
                listener.enterNotEqualExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNotEqualExpr" ):
                listener.exitNotEqualExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNotEqualExpr" ):
                return visitor.visitNotEqualExpr(self)
            else:
                return visitor.visitChildren(self)


    class SelfSubLeftExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExprPtrParserContext
            self.copyFrom(ctx)

        def Sub(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Sub)
            else:
                return self.getToken(RiddleParser.Sub, i)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSelfSubLeftExpr" ):
                listener.enterSelfSubLeftExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSelfSubLeftExpr" ):
                listener.exitSelfSubLeftExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSelfSubLeftExpr" ):
                return visitor.visitSelfSubLeftExpr(self)
            else:
                return visitor.visitChildren(self)


    class DivExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Div(self):
            return self.getToken(RiddleParser.Div, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterDivExpr" ):
                listener.enterDivExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitDivExpr" ):
                listener.exitDivExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitDivExpr" ):
                return visitor.visitDivExpr(self)
            else:
                return visitor.visitChildren(self)


    class BitAndExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def And(self):
            return self.getToken(RiddleParser.And, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBitAndExpr" ):
                listener.enterBitAndExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBitAndExpr" ):
                listener.exitBitAndExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBitAndExpr" ):
                return visitor.visitBitAndExpr(self)
            else:
                return visitor.visitChildren(self)


    class AssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAssignExpr" ):
                listener.enterAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAssignExpr" ):
                listener.exitAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAssignExpr" ):
                return visitor.visitAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class SelfAddRightExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExprPtrParserContext
            self.copyFrom(ctx)

        def Add(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Add)
            else:
                return self.getToken(RiddleParser.Add, i)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSelfAddRightExpr" ):
                listener.enterSelfAddRightExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSelfAddRightExpr" ):
                listener.exitSelfAddRightExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSelfAddRightExpr" ):
                return visitor.visitSelfAddRightExpr(self)
            else:
                return visitor.visitChildren(self)


    class SubAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Sub(self):
            return self.getToken(RiddleParser.Sub, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSubAssignExpr" ):
                listener.enterSubAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSubAssignExpr" ):
                listener.exitSubAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSubAssignExpr" ):
                return visitor.visitSubAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class BracketExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExpressionContext
            self.copyFrom(ctx)

        def LeftBracket(self):
            return self.getToken(RiddleParser.LeftBracket, 0)
        def RightBracket(self):
            return self.getToken(RiddleParser.RightBracket, 0)
        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBracketExpr" ):
                listener.enterBracketExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBracketExpr" ):
                listener.exitBracketExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBracketExpr" ):
                return visitor.visitBracketExpr(self)
            else:
                return visitor.visitChildren(self)


    class BooleanExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.copyFrom(ctx)

        def boolean(self):
            return self.getTypedRuleContext(RiddleParser.BooleanContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBooleanExpr" ):
                listener.enterBooleanExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBooleanExpr" ):
                listener.exitBooleanExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBooleanExpr" ):
                return visitor.visitBooleanExpr(self)
            else:
                return visitor.visitChildren(self)


    class AShrExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def RightRight(self):
            return self.getToken(RiddleParser.RightRight, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAShrExpr" ):
                listener.enterAShrExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAShrExpr" ):
                listener.exitAShrExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAShrExpr" ):
                return visitor.visitAShrExpr(self)
            else:
                return visitor.visitChildren(self)


    class LessExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Less(self):
            return self.getToken(RiddleParser.Less, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLessExpr" ):
                listener.enterLessExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLessExpr" ):
                listener.exitLessExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLessExpr" ):
                return visitor.visitLessExpr(self)
            else:
                return visitor.visitChildren(self)


    class OrExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Or(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Or)
            else:
                return self.getToken(RiddleParser.Or, i)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterOrExpr" ):
                listener.enterOrExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitOrExpr" ):
                listener.exitOrExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitOrExpr" ):
                return visitor.visitOrExpr(self)
            else:
                return visitor.visitChildren(self)


    class SelfAddLeftExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExprPtrParserContext
            self.copyFrom(ctx)

        def Add(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Add)
            else:
                return self.getToken(RiddleParser.Add, i)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterSelfAddLeftExpr" ):
                listener.enterSelfAddLeftExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitSelfAddLeftExpr" ):
                listener.exitSelfAddLeftExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitSelfAddLeftExpr" ):
                return visitor.visitSelfAddLeftExpr(self)
            else:
                return visitor.visitChildren(self)


    class LShrAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def RightRightRight(self):
            return self.getToken(RiddleParser.RightRightRight, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLShrAssignExpr" ):
                listener.enterLShrAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLShrAssignExpr" ):
                listener.exitLShrAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLShrAssignExpr" ):
                return visitor.visitLShrAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class LessEqualExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Less(self):
            return self.getToken(RiddleParser.Less, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterLessEqualExpr" ):
                listener.enterLessEqualExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitLessEqualExpr" ):
                listener.exitLessEqualExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitLessEqualExpr" ):
                return visitor.visitLessEqualExpr(self)
            else:
                return visitor.visitChildren(self)


    class AShrAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def RightRight(self):
            return self.getToken(RiddleParser.RightRight, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAShrAssignExpr" ):
                listener.enterAShrAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAShrAssignExpr" ):
                listener.exitAShrAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAShrAssignExpr" ):
                return visitor.visitAShrAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class BitOrExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Or(self):
            return self.getToken(RiddleParser.Or, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBitOrExpr" ):
                listener.enterBitOrExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBitOrExpr" ):
                listener.exitBitOrExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBitOrExpr" ):
                return visitor.visitBitOrExpr(self)
            else:
                return visitor.visitChildren(self)


    class NotExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExpressionContext
            self.copyFrom(ctx)

        def Not(self):
            return self.getToken(RiddleParser.Not, 0)
        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNotExpr" ):
                listener.enterNotExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNotExpr" ):
                listener.exitNotExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNotExpr" ):
                return visitor.visitNotExpr(self)
            else:
                return visitor.visitChildren(self)


    class AddExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Add(self):
            return self.getToken(RiddleParser.Add, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAddExpr" ):
                listener.enterAddExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAddExpr" ):
                listener.exitAddExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAddExpr" ):
                return visitor.visitAddExpr(self)
            else:
                return visitor.visitChildren(self)


    class ShlAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def LeftLeft(self):
            return self.getToken(RiddleParser.LeftLeft, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterShlAssignExpr" ):
                listener.enterShlAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitShlAssignExpr" ):
                listener.exitShlAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitShlAssignExpr" ):
                return visitor.visitShlAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class ModAssignExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExprPtrParserContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Mod(self):
            return self.getToken(RiddleParser.Mod, 0)
        def Assign(self):
            return self.getToken(RiddleParser.Assign, 0)
        def exprPtrParser(self):
            return self.getTypedRuleContext(RiddleParser.ExprPtrParserContext,0)

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterModAssignExpr" ):
                listener.enterModAssignExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitModAssignExpr" ):
                listener.exitModAssignExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitModAssignExpr" ):
                return visitor.visitModAssignExpr(self)
            else:
                return visitor.visitChildren(self)


    class PositiveExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.expr = None # ExpressionContext
            self.copyFrom(ctx)

        def Add(self):
            return self.getToken(RiddleParser.Add, 0)
        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterPositiveExpr" ):
                listener.enterPositiveExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitPositiveExpr" ):
                listener.exitPositiveExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitPositiveExpr" ):
                return visitor.visitPositiveExpr(self)
            else:
                return visitor.visitChildren(self)


    class ShlExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def LeftLeft(self):
            return self.getToken(RiddleParser.LeftLeft, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterShlExpr" ):
                listener.enterShlExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitShlExpr" ):
                listener.exitShlExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitShlExpr" ):
                return visitor.visitShlExpr(self)
            else:
                return visitor.visitChildren(self)


    class EqualExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def Equal(self):
            return self.getToken(RiddleParser.Equal, 0)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterEqualExpr" ):
                listener.enterEqualExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitEqualExpr" ):
                listener.exitEqualExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitEqualExpr" ):
                return visitor.visitEqualExpr(self)
            else:
                return visitor.visitChildren(self)


    class AndExprContext(ExpressionContext):

        def __init__(self, parser, ctx:ParserRuleContext): # actually a RiddleParser.ExpressionContext
            super().__init__(parser)
            self.left = None # ExpressionContext
            self.right = None # ExpressionContext
            self.copyFrom(ctx)

        def And(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.And)
            else:
                return self.getToken(RiddleParser.And, i)
        def expression(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.ExpressionContext)
            else:
                return self.getTypedRuleContext(RiddleParser.ExpressionContext,i)


        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterAndExpr" ):
                listener.enterAndExpr(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitAndExpr" ):
                listener.exitAndExpr(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitAndExpr" ):
                return visitor.visitAndExpr(self)
            else:
                return visitor.visitChildren(self)



    def expression(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = RiddleParser.ExpressionContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 40
        self.enterRecursionRule(localctx, 40, self.RULE_expression, _p)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 356
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,22,self._ctx)
            if la_ == 1:
                localctx = RiddleParser.CastExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx

                self.state = 256
                self.match(RiddleParser.Less)
                self.state = 257
                localctx.type_ = self.typeName(0)
                self.state = 258
                self.match(RiddleParser.Greater)
                self.state = 259
                self.match(RiddleParser.LeftBracket)
                self.state = 260
                localctx.value = self.exprPtrParser()
                self.state = 261
                self.match(RiddleParser.RightBracket)
                pass

            elif la_ == 2:
                localctx = RiddleParser.BracketExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 263
                self.match(RiddleParser.LeftBracket)
                self.state = 264
                localctx.expr = self.expression(0)
                self.state = 265
                self.match(RiddleParser.RightBracket)
                pass

            elif la_ == 3:
                localctx = RiddleParser.NotExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 267
                self.match(RiddleParser.Not)
                self.state = 268
                localctx.expr = self.expression(44)
                pass

            elif la_ == 4:
                localctx = RiddleParser.PositiveExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 269
                self.match(RiddleParser.Add)
                self.state = 270
                localctx.expr = self.expression(43)
                pass

            elif la_ == 5:
                localctx = RiddleParser.NegativeExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 271
                self.match(RiddleParser.Sub)
                self.state = 272
                localctx.expr = self.expression(42)
                pass

            elif la_ == 6:
                localctx = RiddleParser.SelfAddLeftExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 273
                self.match(RiddleParser.Add)
                self.state = 274
                self.match(RiddleParser.Add)
                self.state = 275
                localctx.expr = self.exprPtrParser()
                pass

            elif la_ == 7:
                localctx = RiddleParser.SelfAddRightExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 276
                localctx.expr = self.exprPtrParser()
                self.state = 277
                self.match(RiddleParser.Add)
                self.state = 278
                self.match(RiddleParser.Add)
                pass

            elif la_ == 8:
                localctx = RiddleParser.SelfSubLeftExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 280
                self.match(RiddleParser.Sub)
                self.state = 281
                self.match(RiddleParser.Sub)
                self.state = 282
                localctx.expr = self.exprPtrParser()
                pass

            elif la_ == 9:
                localctx = RiddleParser.SelfSubRightExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 283
                localctx.expr = self.exprPtrParser()
                self.state = 284
                self.match(RiddleParser.Sub)
                self.state = 285
                self.match(RiddleParser.Sub)
                pass

            elif la_ == 10:
                localctx = RiddleParser.PtrExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 287
                self.exprPtr(0)
                pass

            elif la_ == 11:
                localctx = RiddleParser.AssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 288
                localctx.left = self.exprPtrParser()
                self.state = 289
                self.match(RiddleParser.Assign)
                self.state = 290
                localctx.right = self.expression(17)
                pass

            elif la_ == 12:
                localctx = RiddleParser.AddAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 292
                localctx.left = self.exprPtrParser()
                self.state = 293
                self.match(RiddleParser.Add)
                self.state = 294
                self.match(RiddleParser.Assign)
                self.state = 295
                localctx.right = self.expression(16)
                pass

            elif la_ == 13:
                localctx = RiddleParser.SubAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 297
                localctx.left = self.exprPtrParser()
                self.state = 298
                self.match(RiddleParser.Sub)
                self.state = 299
                self.match(RiddleParser.Assign)
                self.state = 300
                localctx.right = self.expression(15)
                pass

            elif la_ == 14:
                localctx = RiddleParser.MulAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 302
                localctx.left = self.exprPtrParser()
                self.state = 303
                self.match(RiddleParser.Star)
                self.state = 304
                self.match(RiddleParser.Assign)
                self.state = 305
                localctx.right = self.expression(14)
                pass

            elif la_ == 15:
                localctx = RiddleParser.DivAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 307
                localctx.left = self.exprPtrParser()
                self.state = 308
                self.match(RiddleParser.Div)
                self.state = 309
                self.match(RiddleParser.Assign)
                self.state = 310
                localctx.right = self.expression(13)
                pass

            elif la_ == 16:
                localctx = RiddleParser.ModAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 312
                localctx.left = self.exprPtrParser()
                self.state = 313
                self.match(RiddleParser.Mod)
                self.state = 314
                self.match(RiddleParser.Assign)
                self.state = 315
                localctx.right = self.expression(12)
                pass

            elif la_ == 17:
                localctx = RiddleParser.AddAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 317
                localctx.left = self.exprPtrParser()
                self.state = 318
                self.match(RiddleParser.Add)
                self.state = 319
                self.match(RiddleParser.Assign)
                self.state = 320
                localctx.right = self.expression(11)
                pass

            elif la_ == 18:
                localctx = RiddleParser.AndAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 322
                localctx.left = self.exprPtrParser()
                self.state = 323
                self.match(RiddleParser.And)
                self.state = 324
                self.match(RiddleParser.Assign)
                self.state = 325
                localctx.right = self.expression(10)
                pass

            elif la_ == 19:
                localctx = RiddleParser.OrAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 327
                localctx.left = self.exprPtrParser()
                self.state = 328
                self.match(RiddleParser.Or)
                self.state = 329
                self.match(RiddleParser.Assign)
                self.state = 330
                localctx.right = self.expression(9)
                pass

            elif la_ == 20:
                localctx = RiddleParser.XorAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 332
                localctx.left = self.exprPtrParser()
                self.state = 333
                self.match(RiddleParser.Xor)
                self.state = 334
                self.match(RiddleParser.Assign)
                self.state = 335
                localctx.right = self.expression(8)
                pass

            elif la_ == 21:
                localctx = RiddleParser.ShlAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 337
                localctx.left = self.exprPtrParser()
                self.state = 338
                self.match(RiddleParser.LeftLeft)
                self.state = 339
                self.match(RiddleParser.Assign)
                self.state = 340
                localctx.right = self.expression(7)
                pass

            elif la_ == 22:
                localctx = RiddleParser.AShrAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 342
                localctx.left = self.exprPtrParser()
                self.state = 343
                self.match(RiddleParser.RightRight)
                self.state = 344
                self.match(RiddleParser.Assign)
                self.state = 345
                localctx.right = self.expression(6)
                pass

            elif la_ == 23:
                localctx = RiddleParser.LShrAssignExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 347
                localctx.left = self.exprPtrParser()
                self.state = 348
                self.match(RiddleParser.RightRightRight)
                self.state = 349
                self.match(RiddleParser.Assign)
                self.state = 350
                localctx.right = self.expression(5)
                pass

            elif la_ == 24:
                localctx = RiddleParser.StringExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 352
                self.string()
                pass

            elif la_ == 25:
                localctx = RiddleParser.NumberExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 353
                self.number()
                pass

            elif la_ == 26:
                localctx = RiddleParser.BooleanExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 354
                self.boolean()
                pass

            elif la_ == 27:
                localctx = RiddleParser.NullExprContext(self, localctx)
                self._ctx = localctx
                _prevctx = localctx
                self.state = 355
                self.match(RiddleParser.Null)
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 422
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,24,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    self.state = 420
                    self._errHandler.sync(self)
                    la_ = self._interp.adaptivePredict(self._input,23,self._ctx)
                    if la_ == 1:
                        localctx = RiddleParser.MulExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 358
                        if not self.precpred(self._ctx, 36):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 36)")
                        self.state = 359
                        self.match(RiddleParser.Star)
                        self.state = 360
                        localctx.right = self.expression(37)
                        pass

                    elif la_ == 2:
                        localctx = RiddleParser.DivExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 361
                        if not self.precpred(self._ctx, 35):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 35)")
                        self.state = 362
                        self.match(RiddleParser.Div)
                        self.state = 363
                        localctx.right = self.expression(36)
                        pass

                    elif la_ == 3:
                        localctx = RiddleParser.ModExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 364
                        if not self.precpred(self._ctx, 34):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 34)")
                        self.state = 365
                        self.match(RiddleParser.Mod)
                        self.state = 366
                        localctx.right = self.expression(35)
                        pass

                    elif la_ == 4:
                        localctx = RiddleParser.AddExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 367
                        if not self.precpred(self._ctx, 33):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 33)")
                        self.state = 368
                        self.match(RiddleParser.Add)
                        self.state = 369
                        localctx.right = self.expression(34)
                        pass

                    elif la_ == 5:
                        localctx = RiddleParser.SubExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 370
                        if not self.precpred(self._ctx, 32):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 32)")
                        self.state = 371
                        self.match(RiddleParser.Sub)
                        self.state = 372
                        localctx.right = self.expression(33)
                        pass

                    elif la_ == 6:
                        localctx = RiddleParser.ShlExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 373
                        if not self.precpred(self._ctx, 31):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 31)")
                        self.state = 374
                        self.match(RiddleParser.LeftLeft)
                        self.state = 375
                        localctx.right = self.expression(32)
                        pass

                    elif la_ == 7:
                        localctx = RiddleParser.AShrExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 376
                        if not self.precpred(self._ctx, 30):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 30)")
                        self.state = 377
                        self.match(RiddleParser.RightRight)
                        self.state = 378
                        localctx.right = self.expression(31)
                        pass

                    elif la_ == 8:
                        localctx = RiddleParser.LShrExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 379
                        if not self.precpred(self._ctx, 29):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 29)")
                        self.state = 380
                        self.match(RiddleParser.RightRightRight)
                        self.state = 381
                        localctx.right = self.expression(30)
                        pass

                    elif la_ == 9:
                        localctx = RiddleParser.GreaterExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 382
                        if not self.precpred(self._ctx, 28):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 28)")
                        self.state = 383
                        self.match(RiddleParser.Greater)
                        self.state = 384
                        localctx.right = self.expression(29)
                        pass

                    elif la_ == 10:
                        localctx = RiddleParser.LessExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 385
                        if not self.precpred(self._ctx, 27):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 27)")
                        self.state = 386
                        self.match(RiddleParser.Less)
                        self.state = 387
                        localctx.right = self.expression(28)
                        pass

                    elif la_ == 11:
                        localctx = RiddleParser.GreaterEqualExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 388
                        if not self.precpred(self._ctx, 26):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 26)")
                        self.state = 389
                        self.match(RiddleParser.Greater)
                        self.state = 390
                        self.match(RiddleParser.Assign)
                        self.state = 391
                        localctx.right = self.expression(27)
                        pass

                    elif la_ == 12:
                        localctx = RiddleParser.LessEqualExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 392
                        if not self.precpred(self._ctx, 25):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 25)")
                        self.state = 393
                        self.match(RiddleParser.Less)
                        self.state = 394
                        self.match(RiddleParser.Assign)
                        self.state = 395
                        localctx.right = self.expression(26)
                        pass

                    elif la_ == 13:
                        localctx = RiddleParser.EqualExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 396
                        if not self.precpred(self._ctx, 24):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 24)")
                        self.state = 397
                        self.match(RiddleParser.Equal)
                        self.state = 398
                        localctx.right = self.expression(25)
                        pass

                    elif la_ == 14:
                        localctx = RiddleParser.NotEqualExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 399
                        if not self.precpred(self._ctx, 23):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 23)")
                        self.state = 400
                        self.match(RiddleParser.Not)
                        self.state = 401
                        self.match(RiddleParser.Assign)
                        self.state = 402
                        localctx.right = self.expression(24)
                        pass

                    elif la_ == 15:
                        localctx = RiddleParser.BitAndExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 403
                        if not self.precpred(self._ctx, 22):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 22)")
                        self.state = 404
                        self.match(RiddleParser.And)
                        self.state = 405
                        localctx.right = self.expression(23)
                        pass

                    elif la_ == 16:
                        localctx = RiddleParser.BitXorExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 406
                        if not self.precpred(self._ctx, 21):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 21)")
                        self.state = 407
                        self.match(RiddleParser.Xor)
                        self.state = 408
                        localctx.right = self.expression(22)
                        pass

                    elif la_ == 17:
                        localctx = RiddleParser.BitOrExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 409
                        if not self.precpred(self._ctx, 20):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 20)")
                        self.state = 410
                        self.match(RiddleParser.Or)
                        self.state = 411
                        localctx.right = self.expression(21)
                        pass

                    elif la_ == 18:
                        localctx = RiddleParser.AndExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 412
                        if not self.precpred(self._ctx, 19):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 19)")
                        self.state = 413
                        self.match(RiddleParser.And)
                        self.state = 414
                        self.match(RiddleParser.And)
                        self.state = 415
                        localctx.right = self.expression(20)
                        pass

                    elif la_ == 19:
                        localctx = RiddleParser.OrExprContext(self, RiddleParser.ExpressionContext(self, _parentctx, _parentState))
                        localctx.left = _prevctx
                        self.pushNewRecursionContext(localctx, _startState, self.RULE_expression)
                        self.state = 416
                        if not self.precpred(self._ctx, 18):
                            from antlr4.error.Errors import FailedPredicateException
                            raise FailedPredicateException(self, "self.precpred(self._ctx, 18)")
                        self.state = 417
                        self.match(RiddleParser.Or)
                        self.state = 418
                        self.match(RiddleParser.Or)
                        self.state = 419
                        localctx.right = self.expression(19)
                        pass

             
                self.state = 424
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,24,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx


    class IdContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def Identifier(self):
            return self.getToken(RiddleParser.Identifier, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_id

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterId" ):
                listener.enterId(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitId" ):
                listener.exitId(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitId" ):
                return visitor.visitId(self)
            else:
                return visitor.visitChildren(self)




    def id_(self):

        localctx = RiddleParser.IdContext(self, self._ctx, self.state)
        self.enterRule(localctx, 42, self.RULE_id)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 425
            self.match(RiddleParser.Identifier)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class NumberContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def integer(self):
            return self.getTypedRuleContext(RiddleParser.IntegerContext,0)


        def float_(self):
            return self.getTypedRuleContext(RiddleParser.FloatContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_number

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterNumber" ):
                listener.enterNumber(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitNumber" ):
                listener.exitNumber(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitNumber" ):
                return visitor.visitNumber(self)
            else:
                return visitor.visitChildren(self)




    def number(self):

        localctx = RiddleParser.NumberContext(self, self._ctx, self.state)
        self.enterRule(localctx, 44, self.RULE_number)
        try:
            self.state = 429
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [53, 54, 55, 56]:
                self.enterOuterAlt(localctx, 1)
                self.state = 427
                self.integer()
                pass
            elif token in [57]:
                self.enterOuterAlt(localctx, 2)
                self.state = 428
                self.float_()
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class BooleanContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.value = None

        def True_(self):
            return self.getToken(RiddleParser.True_, 0)

        def False_(self):
            return self.getToken(RiddleParser.False_, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_boolean

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterBoolean" ):
                listener.enterBoolean(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitBoolean" ):
                listener.exitBoolean(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitBoolean" ):
                return visitor.visitBoolean(self)
            else:
                return visitor.visitChildren(self)




    def boolean(self):

        localctx = RiddleParser.BooleanContext(self, self._ctx, self.state)
        self.enterRule(localctx, 46, self.RULE_boolean)
        try:
            self.state = 435
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [16]:
                self.enterOuterAlt(localctx, 1)
                self.state = 431
                self.match(RiddleParser.True_)
                localctx.value=True
                pass
            elif token in [17]:
                self.enterOuterAlt(localctx, 2)
                self.state = 433
                self.match(RiddleParser.False_)
                localctx.value=False
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class StringContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def STRING(self):
            return self.getToken(RiddleParser.STRING, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_string

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterString" ):
                listener.enterString(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitString" ):
                listener.exitString(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitString" ):
                return visitor.visitString(self)
            else:
                return visitor.visitChildren(self)




    def string(self):

        localctx = RiddleParser.StringContext(self, self._ctx, self.state)
        self.enterRule(localctx, 48, self.RULE_string)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 437
            self.match(RiddleParser.STRING)
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class FloatContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.value = None
            self._Float = None # Token

        def Float(self):
            return self.getToken(RiddleParser.Float, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_float

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterFloat" ):
                listener.enterFloat(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitFloat" ):
                listener.exitFloat(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitFloat" ):
                return visitor.visitFloat(self)
            else:
                return visitor.visitChildren(self)




    def float_(self):

        localctx = RiddleParser.FloatContext(self, self._ctx, self.state)
        self.enterRule(localctx, 50, self.RULE_float)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 439
            localctx._Float = self.match(RiddleParser.Float)
            localctx.value = float((None if localctx._Float is None else localctx._Float.text))
        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class IntegerContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.value = None
            self._Decimal = None # Token
            self._Hexadecimal = None # Token
            self._Binary = None # Token
            self._Octal = None # Token

        def Decimal(self):
            return self.getToken(RiddleParser.Decimal, 0)

        def Hexadecimal(self):
            return self.getToken(RiddleParser.Hexadecimal, 0)

        def Binary(self):
            return self.getToken(RiddleParser.Binary, 0)

        def Octal(self):
            return self.getToken(RiddleParser.Octal, 0)

        def getRuleIndex(self):
            return RiddleParser.RULE_integer

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterInteger" ):
                listener.enterInteger(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitInteger" ):
                listener.exitInteger(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitInteger" ):
                return visitor.visitInteger(self)
            else:
                return visitor.visitChildren(self)




    def integer(self):

        localctx = RiddleParser.IntegerContext(self, self._ctx, self.state)
        self.enterRule(localctx, 52, self.RULE_integer)
        try:
            self.state = 450
            self._errHandler.sync(self)
            token = self._input.LA(1)
            if token in [54]:
                self.enterOuterAlt(localctx, 1)
                self.state = 442
                localctx._Decimal = self.match(RiddleParser.Decimal)
                localctx.value = int((None if localctx._Decimal is None else localctx._Decimal.text))
                pass
            elif token in [53]:
                self.enterOuterAlt(localctx, 2)
                self.state = 444
                localctx._Hexadecimal = self.match(RiddleParser.Hexadecimal)
                localctx.value = int((None if localctx._Hexadecimal is None else localctx._Hexadecimal.text)[2:],16)
                pass
            elif token in [56]:
                self.enterOuterAlt(localctx, 3)
                self.state = 446
                localctx._Binary = self.match(RiddleParser.Binary)
                localctx.value = int((None if localctx._Binary is None else localctx._Binary.text)[2:],2)
                pass
            elif token in [55]:
                self.enterOuterAlt(localctx, 4)
                self.state = 448
                localctx._Octal = self.match(RiddleParser.Octal)
                localctx.value = int((None if localctx._Octal is None else localctx._Octal.text)[1:],8)
                pass
            else:
                raise NoViableAltException(self)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TemplateArgContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def typeName(self):
            return self.getTypedRuleContext(RiddleParser.TypeNameContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_templateArg

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTemplateArg" ):
                listener.enterTemplateArg(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTemplateArg" ):
                listener.exitTemplateArg(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitTemplateArg" ):
                return visitor.visitTemplateArg(self)
            else:
                return visitor.visitChildren(self)




    def templateArg(self):

        localctx = RiddleParser.TemplateArgContext(self, self._ctx, self.state)
        self.enterRule(localctx, 54, self.RULE_templateArg)
        try:
            self.state = 454
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,28,self._ctx)
            if la_ == 1:
                self.enterOuterAlt(localctx, 1)
                self.state = 452
                self.expression(0)
                pass

            elif la_ == 2:
                self.enterOuterAlt(localctx, 2)
                self.state = 453
                self.typeName(0)
                pass


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TemplateArgsContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser

        def templateArg(self, i:int=None):
            if i is None:
                return self.getTypedRuleContexts(RiddleParser.TemplateArgContext)
            else:
                return self.getTypedRuleContext(RiddleParser.TemplateArgContext,i)


        def Comma(self, i:int=None):
            if i is None:
                return self.getTokens(RiddleParser.Comma)
            else:
                return self.getToken(RiddleParser.Comma, i)

        def getRuleIndex(self):
            return RiddleParser.RULE_templateArgs

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTemplateArgs" ):
                listener.enterTemplateArgs(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTemplateArgs" ):
                listener.exitTemplateArgs(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitTemplateArgs" ):
                return visitor.visitTemplateArgs(self)
            else:
                return visitor.visitChildren(self)




    def templateArgs(self):

        localctx = RiddleParser.TemplateArgsContext(self, self._ctx, self.state)
        self.enterRule(localctx, 56, self.RULE_templateArgs)
        self._la = 0 # Token type
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 465
            self._errHandler.sync(self)
            _la = self._input.LA(1)
            if (((_la) & ~0x3f) == 0 and ((1 << _la) & -8939625984507576320) != 0):
                self.state = 461
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,29,self._ctx)
                while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                    if _alt==1:
                        self.state = 456
                        self.templateArg()
                        self.state = 457
                        self.match(RiddleParser.Comma) 
                    self.state = 463
                    self._errHandler.sync(self)
                    _alt = self._interp.adaptivePredict(self._input,29,self._ctx)

                self.state = 464
                self.templateArg()


        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.exitRule()
        return localctx


    class TypeNameContext(ParserRuleContext):
        __slots__ = 'parser'

        def __init__(self, parser, parent:ParserRuleContext=None, invokingState:int=-1):
            super().__init__(parent, invokingState)
            self.parser = parser
            self.baseType = None # TypeNameContext
            self.name = None # IdContext
            self.args = None # TemplateArgsContext
            self.size = None # ExpressionContext

        def id_(self):
            return self.getTypedRuleContext(RiddleParser.IdContext,0)


        def Less(self):
            return self.getToken(RiddleParser.Less, 0)

        def Greater(self):
            return self.getToken(RiddleParser.Greater, 0)

        def templateArgs(self):
            return self.getTypedRuleContext(RiddleParser.TemplateArgsContext,0)


        def LeftSquare(self):
            return self.getToken(RiddleParser.LeftSquare, 0)

        def RightSquare(self):
            return self.getToken(RiddleParser.RightSquare, 0)

        def typeName(self):
            return self.getTypedRuleContext(RiddleParser.TypeNameContext,0)


        def expression(self):
            return self.getTypedRuleContext(RiddleParser.ExpressionContext,0)


        def getRuleIndex(self):
            return RiddleParser.RULE_typeName

        def enterRule(self, listener:ParseTreeListener):
            if hasattr( listener, "enterTypeName" ):
                listener.enterTypeName(self)

        def exitRule(self, listener:ParseTreeListener):
            if hasattr( listener, "exitTypeName" ):
                listener.exitTypeName(self)

        def accept(self, visitor:ParseTreeVisitor):
            if hasattr( visitor, "visitTypeName" ):
                return visitor.visitTypeName(self)
            else:
                return visitor.visitChildren(self)



    def typeName(self, _p:int=0):
        _parentctx = self._ctx
        _parentState = self.state
        localctx = RiddleParser.TypeNameContext(self, self._ctx, _parentState)
        _prevctx = localctx
        _startState = 58
        self.enterRecursionRule(localctx, 58, self.RULE_typeName, _p)
        try:
            self.enterOuterAlt(localctx, 1)
            self.state = 474
            self._errHandler.sync(self)
            la_ = self._interp.adaptivePredict(self._input,31,self._ctx)
            if la_ == 1:
                self.state = 468
                localctx.name = self.id_()
                pass

            elif la_ == 2:
                self.state = 469
                localctx.name = self.id_()
                self.state = 470
                self.match(RiddleParser.Less)
                self.state = 471
                localctx.args = self.templateArgs()
                self.state = 472
                self.match(RiddleParser.Greater)
                pass


            self._ctx.stop = self._input.LT(-1)
            self.state = 483
            self._errHandler.sync(self)
            _alt = self._interp.adaptivePredict(self._input,32,self._ctx)
            while _alt!=2 and _alt!=ATN.INVALID_ALT_NUMBER:
                if _alt==1:
                    if self._parseListeners is not None:
                        self.triggerExitRuleEvent()
                    _prevctx = localctx
                    localctx = RiddleParser.TypeNameContext(self, _parentctx, _parentState)
                    localctx.baseType = _prevctx
                    self.pushNewRecursionContext(localctx, _startState, self.RULE_typeName)
                    self.state = 476
                    if not self.precpred(self._ctx, 1):
                        from antlr4.error.Errors import FailedPredicateException
                        raise FailedPredicateException(self, "self.precpred(self._ctx, 1)")
                    self.state = 477
                    self.match(RiddleParser.LeftSquare)
                    self.state = 478
                    localctx.size = self.expression(0)
                    self.state = 479
                    self.match(RiddleParser.RightSquare) 
                self.state = 485
                self._errHandler.sync(self)
                _alt = self._interp.adaptivePredict(self._input,32,self._ctx)

        except RecognitionException as re:
            localctx.exception = re
            self._errHandler.reportError(self, re)
            self._errHandler.recover(self, re)
        finally:
            self.unrollRecursionContexts(_parentctx)
        return localctx



    def sempred(self, localctx:RuleContext, ruleIndex:int, predIndex:int):
        if self._predicates == None:
            self._predicates = dict()
        self._predicates[18] = self.exprPtr_sempred
        self._predicates[20] = self.expression_sempred
        self._predicates[29] = self.typeName_sempred
        pred = self._predicates.get(ruleIndex, None)
        if pred is None:
            raise Exception("No predicate with index:" + str(ruleIndex))
        else:
            return pred(localctx, predIndex)

    def exprPtr_sempred(self, localctx:ExprPtrContext, predIndex:int):
            if predIndex == 0:
                return self.precpred(self._ctx, 1)
         

    def expression_sempred(self, localctx:ExpressionContext, predIndex:int):
            if predIndex == 1:
                return self.precpred(self._ctx, 36)
         

            if predIndex == 2:
                return self.precpred(self._ctx, 35)
         

            if predIndex == 3:
                return self.precpred(self._ctx, 34)
         

            if predIndex == 4:
                return self.precpred(self._ctx, 33)
         

            if predIndex == 5:
                return self.precpred(self._ctx, 32)
         

            if predIndex == 6:
                return self.precpred(self._ctx, 31)
         

            if predIndex == 7:
                return self.precpred(self._ctx, 30)
         

            if predIndex == 8:
                return self.precpred(self._ctx, 29)
         

            if predIndex == 9:
                return self.precpred(self._ctx, 28)
         

            if predIndex == 10:
                return self.precpred(self._ctx, 27)
         

            if predIndex == 11:
                return self.precpred(self._ctx, 26)
         

            if predIndex == 12:
                return self.precpred(self._ctx, 25)
         

            if predIndex == 13:
                return self.precpred(self._ctx, 24)
         

            if predIndex == 14:
                return self.precpred(self._ctx, 23)
         

            if predIndex == 15:
                return self.precpred(self._ctx, 22)
         

            if predIndex == 16:
                return self.precpred(self._ctx, 21)
         

            if predIndex == 17:
                return self.precpred(self._ctx, 20)
         

            if predIndex == 18:
                return self.precpred(self._ctx, 19)
         

            if predIndex == 19:
                return self.precpred(self._ctx, 18)
         

    def typeName_sempred(self, localctx:TypeNameContext, predIndex:int):
            if predIndex == 20:
                return self.precpred(self._ctx, 1)
         




