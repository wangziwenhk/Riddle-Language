
// Generated from C:/Users/wangz/CLionProjects/Riddle-Language/RiddleLexer.g4 by ANTLR 4.13.1


#include "RiddleLexer.h"


using namespace antlr4;



using namespace antlr4;

namespace {

struct RiddleLexerStaticData final {
  RiddleLexerStaticData(std::vector<std::string> ruleNames,
                          std::vector<std::string> channelNames,
                          std::vector<std::string> modeNames,
                          std::vector<std::string> literalNames,
                          std::vector<std::string> symbolicNames)
      : ruleNames(std::move(ruleNames)), channelNames(std::move(channelNames)),
        modeNames(std::move(modeNames)), literalNames(std::move(literalNames)),
        symbolicNames(std::move(symbolicNames)),
        vocabulary(this->literalNames, this->symbolicNames) {}

  RiddleLexerStaticData(const RiddleLexerStaticData&) = delete;
  RiddleLexerStaticData(RiddleLexerStaticData&&) = delete;
  RiddleLexerStaticData& operator=(const RiddleLexerStaticData&) = delete;
  RiddleLexerStaticData& operator=(RiddleLexerStaticData&&) = delete;

  std::vector<antlr4::dfa::DFA> decisionToDFA;
  antlr4::atn::PredictionContextCache sharedContextCache;
  const std::vector<std::string> ruleNames;
  const std::vector<std::string> channelNames;
  const std::vector<std::string> modeNames;
  const std::vector<std::string> literalNames;
  const std::vector<std::string> symbolicNames;
  const antlr4::dfa::Vocabulary vocabulary;
  antlr4::atn::SerializedATNView serializedATN;
  std::unique_ptr<antlr4::atn::ATN> atn;
};

::antlr4::internal::OnceFlag riddlelexerLexerOnceFlag;
#if ANTLR4_USE_THREAD_LOCAL_CACHE
static thread_local
#endif
RiddleLexerStaticData *riddlelexerLexerStaticData = nullptr;

void riddlelexerLexerInitialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  if (riddlelexerLexerStaticData != nullptr) {
    return;
  }
#else
  assert(riddlelexerLexerStaticData == nullptr);
#endif
  auto staticData = std::make_unique<RiddleLexerStaticData>(
    std::vector<std::string>{
                  "Var", "Val", "For", "While", "If", "Else", "Func", "Return", "Import",
                  "Package", "Class", "Public", "Protected", "Private", "Override",
                  "True", "False", "LeftBracket", "RightBracket", "LeftSquare", "RightSquare",
                  "LeftCurly", "RightCurly", "Colon", "Semi", "Comma", "Equal", "Assign",
                  "Greater", "Less", "LeftLeft", "RightRight", "RightRightRight", "Add",
                  "Sub", "Star", "Div", "Mod", "Not", "And", "Or", "Xor", "Dot", "DoubleQuotes",
                  "Quotes", "Endl", "Identifier", "Hexadecimal", "Decimal", "Octal",
                  "Binary", "Float", "IntegerSequence", "HEX_DIGIT", "OCTAL_DIGIT",
                  "BINARY_DIGIT", "DIGIT", "STRING", "ESC", "LINE_COMMENT", "BLOCK_COMMENT",
                  "WHITESPACE"},
          std::vector<std::string>{
                  "DEFAULT_TOKEN_CHANNEL", "HIDDEN"},
          std::vector<std::string>{
      "DEFAULT_MODE"
    },
    std::vector<std::string>{
                  "", "'var'", "'val'", "'for'", "'while'", "'if'", "'else'", "'fun'",
                  "'return'", "'import'", "'package'", "'class'", "'public'", "'protected'",
                  "'Private'", "'override'", "'true'", "'false'", "'('", "')'", "'['",
                  "']'", "'{'", "'}'", "':'", "';'", "','", "'=='", "'='", "'>'", "'<'",
                  "'<<'", "'>>'", "'>>>'", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'",
                  "'&'", "'|'", "'^'", "'.'", "'\"'", "'''", "'\\n'"},
          std::vector<std::string>{
                  "", "Var", "Val", "For", "While", "If", "Else", "Func", "Return",
                  "Import", "Package", "Class", "Public", "Protected", "Private", "Override",
                  "True", "False", "LeftBracket", "RightBracket", "LeftSquare", "RightSquare",
                  "LeftCurly", "RightCurly", "Colon", "Semi", "Comma", "Equal", "Assign",
                  "Greater", "Less", "LeftLeft", "RightRight", "RightRightRight", "Add",
                  "Sub", "Star", "Div", "Mod", "Not", "And", "Or", "Xor", "Dot", "DoubleQuotes",
                  "Quotes", "Endl", "Identifier", "Hexadecimal", "Decimal", "Octal",
                  "Binary", "Float", "IntegerSequence", "HEX_DIGIT", "OCTAL_DIGIT",
                  "BINARY_DIGIT", "DIGIT", "STRING", "LINE_COMMENT", "BLOCK_COMMENT",
                  "WHITESPACE"});
  static const int32_t serializedATNSegment[] = {
          4, 0, 61, 399, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4, 2, 5, 7, 5, 2, 6, 7,
          6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14,
          7, 14, 2, 15, 7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2, 21,
          7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26, 7, 26, 2, 27, 7, 27, 2, 28,
          7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35,
          7,35,2,36,7,36,2,37,7,37,2,38,7,38,2,39,7,39,2,40,7,40,2,41,7,41,2,42,
  	7,42,2,43,7,43,2,44,7,44,2,45,7,45,2,46,7,46,2,47,7,47,2,48,7,48,2,49,
  	7,49,2,50,7,50,2,51,7,51,2,52,7,52,2,53,7,53,2,54,7,54,2,55,7,55,2,56,
          7, 56, 2, 57, 7, 57, 2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 1, 0, 1, 0, 1, 0, 1,
          0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
          1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1,
          8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10,
          1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
          1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
          1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15,
          1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20,
          1, 21, 1, 21, 1, 22, 1, 22, 1, 23, 1, 23, 1, 24, 1, 24, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 27,
          1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32,
          1, 32, 1, 33, 1, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 36, 1, 36, 1, 37, 1, 37, 1, 38, 1, 38, 1, 39,
          1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46,
          1, 46, 5, 46, 294, 8, 46, 10, 46, 12, 46, 297, 9, 46, 1, 47, 1, 47, 1, 47, 1, 47, 1, 47, 5, 47,
          304, 8, 47, 10, 47, 12, 47, 307, 9, 47, 1, 48, 1, 48, 1, 48, 5, 48, 312, 8, 48, 10, 48, 12, 48,
          315, 9, 48, 3, 48, 317, 8, 48, 1, 49, 1, 49, 1, 49, 5, 49, 322, 8, 49, 10, 49, 12, 49, 325, 9,
          49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 50, 5, 50, 332, 8, 50, 10, 50, 12, 50, 335, 9, 50, 1, 51, 1,
          51, 1, 51, 1, 51, 1, 52, 4, 52, 342, 8, 52, 11, 52, 12, 52, 343, 1, 53, 1, 53, 1, 54, 1, 54, 1,
          55, 1, 55, 1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 5, 57, 357, 8, 57, 10, 57, 12, 57, 360, 9, 57, 1,
          57, 1, 57, 1, 58, 1, 58, 3, 58, 366, 8, 58, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 372, 8, 59, 10,
          59, 12, 59, 375, 9, 59, 1, 59, 1, 59, 1, 60, 1, 60, 1, 60, 1, 60, 5, 60, 383, 8, 60, 10, 60, 12,
          60, 386, 9, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 60, 1, 61, 4, 61, 394, 8, 61, 11, 61, 12, 61, 395,
          1, 61, 1, 61, 1, 384, 0, 62, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21,
          11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22,
          45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67,
          34, 69, 35, 71, 36, 73, 37, 75, 38, 77, 39, 79, 40, 81, 41, 83, 42, 85, 43, 87, 44, 89, 45,
          91, 46, 93, 47, 95, 48, 97, 49, 99, 50, 101, 51, 103, 52, 105, 53, 107, 54, 109, 55, 111,
          56, 113, 57, 115, 58, 117, 0, 119, 59, 121, 60, 123, 61, 1, 0, 13, 3, 0, 65, 90, 95, 95, 97,
          122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 49, 57, 65, 70, 1, 0, 49, 57, 1, 0, 49, 55,
          2, 0, 48, 57, 65, 70, 1, 0, 48, 55, 1, 0, 48, 49, 1, 0, 48, 57, 4, 0, 10, 10, 13, 13, 34, 34, 92,
          92, 5, 0, 34, 34, 47, 47, 65, 90, 92, 92, 97, 122, 2, 0, 10, 10, 13, 13, 3, 0, 9, 10, 13, 13,
          32, 32, 409, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0,
          0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
          1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0,
          0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0,
          0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53,
          1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0,
          0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 0, 69, 1, 0, 0, 0, 0, 71, 1, 0, 0, 0, 0, 73, 1, 0, 0, 0,
          0, 75, 1, 0, 0, 0, 0, 77, 1, 0, 0, 0, 0, 79, 1, 0, 0, 0, 0, 81, 1, 0, 0, 0, 0, 83, 1, 0, 0, 0, 0, 85,
          1, 0, 0, 0, 0, 87, 1, 0, 0, 0, 0, 89, 1, 0, 0, 0, 0, 91, 1, 0, 0, 0, 0, 93, 1, 0, 0, 0, 0, 95, 1, 0,
          0, 0, 0, 97, 1, 0, 0, 0, 0, 99, 1, 0, 0, 0, 0, 101, 1, 0, 0, 0, 0, 103, 1, 0, 0, 0, 0, 105, 1, 0, 0,
          0, 0, 107, 1, 0, 0, 0, 0, 109, 1, 0, 0, 0, 0, 111, 1, 0, 0, 0, 0, 113, 1, 0, 0, 0, 0, 115, 1, 0, 0,
          0, 0, 119, 1, 0, 0, 0, 0, 121, 1, 0, 0, 0, 0, 123, 1, 0, 0, 0, 1, 125, 1, 0, 0, 0, 3, 129, 1, 0, 0,
          0, 5, 133, 1, 0, 0, 0, 7, 137, 1, 0, 0, 0, 9, 143, 1, 0, 0, 0, 11, 146, 1, 0, 0, 0, 13, 151, 1, 0,
          0, 0, 15, 155, 1, 0, 0, 0, 17, 162, 1, 0, 0, 0, 19, 169, 1, 0, 0, 0, 21, 177, 1, 0, 0, 0, 23, 183,
          1, 0, 0, 0, 25, 190, 1, 0, 0, 0, 27, 200, 1, 0, 0, 0, 29, 208, 1, 0, 0, 0, 31, 217, 1, 0, 0, 0, 33,
          222, 1, 0, 0, 0, 35, 228, 1, 0, 0, 0, 37, 230, 1, 0, 0, 0, 39, 232, 1, 0, 0, 0, 41, 234, 1, 0, 0,
          0, 43, 236, 1, 0, 0, 0, 45, 238, 1, 0, 0, 0, 47, 240, 1, 0, 0, 0, 49, 242, 1, 0, 0, 0, 51, 244,
          1, 0, 0, 0, 53, 246, 1, 0, 0, 0, 55, 249, 1, 0, 0, 0, 57, 251, 1, 0, 0, 0, 59, 253, 1, 0, 0, 0, 61,
          255, 1, 0, 0, 0, 63, 258, 1, 0, 0, 0, 65, 261, 1, 0, 0, 0, 67, 265, 1, 0, 0, 0, 69, 267, 1, 0, 0,
          0, 71, 269, 1, 0, 0, 0, 73, 271, 1, 0, 0, 0, 75, 273, 1, 0, 0, 0, 77, 275, 1, 0, 0, 0, 79, 277,
          1, 0, 0, 0, 81, 279, 1, 0, 0, 0, 83, 281, 1, 0, 0, 0, 85, 283, 1, 0, 0, 0, 87, 285, 1, 0, 0, 0, 89,
          287, 1, 0, 0, 0, 91, 289, 1, 0, 0, 0, 93, 291, 1, 0, 0, 0, 95, 298, 1, 0, 0, 0, 97, 316, 1, 0, 0,
          0, 99, 318, 1, 0, 0, 0, 101, 326, 1, 0, 0, 0, 103, 336, 1, 0, 0, 0, 105, 341, 1, 0, 0, 0, 107,
          345, 1, 0, 0, 0, 109, 347, 1, 0, 0, 0, 111, 349, 1, 0, 0, 0, 113, 351, 1, 0, 0, 0, 115, 353, 1,
          0, 0, 0, 117, 363, 1, 0, 0, 0, 119, 367, 1, 0, 0, 0, 121, 378, 1, 0, 0, 0, 123, 393, 1, 0, 0, 0,
          125, 126, 5, 118, 0, 0, 126, 127, 5, 97, 0, 0, 127, 128, 5, 114, 0, 0, 128, 2, 1, 0, 0, 0, 129,
          130, 5, 118, 0, 0, 130, 131, 5, 97, 0, 0, 131, 132, 5, 108, 0, 0, 132, 4, 1, 0, 0, 0, 133, 134,
          5, 102, 0, 0, 134, 135, 5, 111, 0, 0, 135, 136, 5, 114, 0, 0, 136, 6, 1, 0, 0, 0, 137, 138, 5,
          119, 0, 0, 138, 139, 5, 104, 0, 0, 139, 140, 5, 105, 0, 0, 140, 141, 5, 108, 0, 0, 141, 142,
          5, 101, 0, 0, 142, 8, 1, 0, 0, 0, 143, 144, 5, 105, 0, 0, 144, 145, 5, 102, 0, 0, 145, 10, 1,
          0, 0, 0, 146, 147, 5, 101, 0, 0, 147, 148, 5, 108, 0, 0, 148, 149, 5, 115, 0, 0, 149, 150, 5,
          101, 0, 0, 150, 12, 1, 0, 0, 0, 151, 152, 5, 102, 0, 0, 152, 153, 5, 117, 0, 0, 153, 154, 5,
          110, 0, 0, 154, 14, 1, 0, 0, 0, 155, 156, 5, 114, 0, 0, 156, 157, 5, 101, 0, 0, 157, 158, 5,
          116, 0, 0, 158, 159, 5, 117, 0, 0, 159, 160, 5, 114, 0, 0, 160, 161, 5, 110, 0, 0, 161, 16,
          1, 0, 0, 0, 162, 163, 5, 105, 0, 0, 163, 164, 5, 109, 0, 0, 164, 165, 5, 112, 0, 0, 165, 166,
          5, 111, 0, 0, 166, 167, 5, 114, 0, 0, 167, 168, 5, 116, 0, 0, 168, 18, 1, 0, 0, 0, 169, 170,
          5, 112, 0, 0, 170, 171, 5, 97, 0, 0, 171, 172, 5, 99, 0, 0, 172, 173, 5, 107, 0, 0, 173, 174,
          5, 97, 0, 0, 174, 175, 5, 103, 0, 0, 175, 176, 5, 101, 0, 0, 176, 20, 1, 0, 0, 0, 177, 178, 5,
          99, 0, 0, 178, 179, 5, 108, 0, 0, 179, 180, 5, 97, 0, 0, 180, 181, 5, 115, 0, 0, 181, 182, 5,
          115, 0, 0, 182, 22, 1, 0, 0, 0, 183, 184, 5, 112, 0, 0, 184, 185, 5, 117, 0, 0, 185, 186, 5,
          98, 0, 0, 186, 187, 5, 108, 0, 0, 187, 188, 5, 105, 0, 0, 188, 189, 5, 99, 0, 0, 189, 24, 1,
          0, 0, 0, 190, 191, 5, 112, 0, 0, 191, 192, 5, 114, 0, 0, 192, 193, 5, 111, 0, 0, 193, 194, 5,
          116, 0, 0, 194, 195, 5, 101, 0, 0, 195, 196, 5, 99, 0, 0, 196, 197, 5, 116, 0, 0, 197, 198,
          5, 101, 0, 0, 198, 199, 5, 100, 0, 0, 199, 26, 1, 0, 0, 0, 200, 201, 5, 80, 0, 0, 201, 202, 5,
          114, 0, 0, 202, 203, 5, 105, 0, 0, 203, 204, 5, 118, 0, 0, 204, 205, 5, 97, 0, 0, 205, 206,
          5, 116, 0, 0, 206, 207, 5, 101, 0, 0, 207, 28, 1, 0, 0, 0, 208, 209, 5, 111, 0, 0, 209, 210,
          5, 118, 0, 0, 210, 211, 5, 101, 0, 0, 211, 212, 5, 114, 0, 0, 212, 213, 5, 114, 0, 0, 213, 214,
          5, 105, 0, 0, 214, 215, 5, 100, 0, 0, 215, 216, 5, 101, 0, 0, 216, 30, 1, 0, 0, 0, 217, 218,
          5, 116, 0, 0, 218, 219, 5, 114, 0, 0, 219, 220, 5, 117, 0, 0, 220, 221, 5, 101, 0, 0, 221, 32,
          1, 0, 0, 0, 222, 223, 5, 102, 0, 0, 223, 224, 5, 97, 0, 0, 224, 225, 5, 108, 0, 0, 225, 226,
          5, 115, 0, 0, 226, 227, 5, 101, 0, 0, 227, 34, 1, 0, 0, 0, 228, 229, 5, 40, 0, 0, 229, 36, 1,
          0, 0, 0, 230, 231, 5, 41, 0, 0, 231, 38, 1, 0, 0, 0, 232, 233, 5, 91, 0, 0, 233, 40, 1, 0, 0, 0,
          234, 235, 5, 93, 0, 0, 235, 42, 1, 0, 0, 0, 236, 237, 5, 123, 0, 0, 237, 44, 1, 0, 0, 0, 238,
          239, 5, 125, 0, 0, 239, 46, 1, 0, 0, 0, 240, 241, 5, 58, 0, 0, 241, 48, 1, 0, 0, 0, 242, 243,
          5, 59, 0, 0, 243, 50, 1, 0, 0, 0, 244, 245, 5, 44, 0, 0, 245, 52, 1, 0, 0, 0, 246, 247, 5, 61,
          0, 0, 247, 248, 5, 61, 0, 0, 248, 54, 1, 0, 0, 0, 249, 250, 5, 61, 0, 0, 250, 56, 1, 0, 0, 0, 251,
          252, 5, 62, 0, 0, 252, 58, 1, 0, 0, 0, 253, 254, 5, 60, 0, 0, 254, 60, 1, 0, 0, 0, 255, 256, 5,
          60, 0, 0, 256, 257, 5, 60, 0, 0, 257, 62, 1, 0, 0, 0, 258, 259, 5, 62, 0, 0, 259, 260, 5, 62,
          0, 0, 260, 64, 1, 0, 0, 0, 261, 262, 5, 62, 0, 0, 262, 263, 5, 62, 0, 0, 263, 264, 5, 62, 0, 0,
          264, 66, 1, 0, 0, 0, 265, 266, 5, 43, 0, 0, 266, 68, 1, 0, 0, 0, 267, 268, 5, 45, 0, 0, 268, 70,
          1, 0, 0, 0, 269, 270, 5, 42, 0, 0, 270, 72, 1, 0, 0, 0, 271, 272, 5, 47, 0, 0, 272, 74, 1, 0, 0,
          0, 273, 274, 5, 37, 0, 0, 274, 76, 1, 0, 0, 0, 275, 276, 5, 33, 0, 0, 276, 78, 1, 0, 0, 0, 277,
          278, 5, 38, 0, 0, 278, 80, 1, 0, 0, 0, 279, 280, 5, 124, 0, 0, 280, 82, 1, 0, 0, 0, 281, 282,
          5, 94, 0, 0, 282, 84, 1, 0, 0, 0, 283, 284, 5, 46, 0, 0, 284, 86, 1, 0, 0, 0, 285, 286, 5, 34,
          0, 0, 286, 88, 1, 0, 0, 0, 287, 288, 5, 39, 0, 0, 288, 90, 1, 0, 0, 0, 289, 290, 5, 10, 0, 0, 290,
          92, 1, 0, 0, 0, 291, 295, 7, 0, 0, 0, 292, 294, 7, 1, 0, 0, 293, 292, 1, 0, 0, 0, 294, 297, 1,
          0, 0, 0, 295, 293, 1, 0, 0, 0, 295, 296, 1, 0, 0, 0, 296, 94, 1, 0, 0, 0, 297, 295, 1, 0, 0, 0,
          298, 299, 5, 48, 0, 0, 299, 300, 5, 120, 0, 0, 300, 301, 1, 0, 0, 0, 301, 305, 7, 2, 0, 0, 302,
          304, 3, 107, 53, 0, 303, 302, 1, 0, 0, 0, 304, 307, 1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 306,
          1, 0, 0, 0, 306, 96, 1, 0, 0, 0, 307, 305, 1, 0, 0, 0, 308, 317, 5, 48, 0, 0, 309, 313, 7, 3, 0,
          0, 310, 312, 3, 113, 56, 0, 311, 310, 1, 0, 0, 0, 312, 315, 1, 0, 0, 0, 313, 311, 1, 0, 0, 0,
          313, 314, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 316, 308, 1, 0, 0, 0, 316, 309,
          1, 0, 0, 0, 317, 98, 1, 0, 0, 0, 318, 319, 5, 48, 0, 0, 319, 323, 7, 4, 0, 0, 320, 322, 3, 109,
          54, 0, 321, 320, 1, 0, 0, 0, 322, 325, 1, 0, 0, 0, 323, 321, 1, 0, 0, 0, 323, 324, 1, 0, 0, 0,
          324, 100, 1, 0, 0, 0, 325, 323, 1, 0, 0, 0, 326, 327, 5, 48, 0, 0, 327, 328, 5, 98, 0, 0, 328,
          329, 1, 0, 0, 0, 329, 333, 5, 49, 0, 0, 330, 332, 3, 111, 55, 0, 331, 330, 1, 0, 0, 0, 332, 335,
          1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 333, 334, 1, 0, 0, 0, 334, 102, 1, 0, 0, 0, 335, 333, 1, 0, 0,
          0, 336, 337, 3, 97, 48, 0, 337, 338, 3, 85, 42, 0, 338, 339, 3, 105, 52, 0, 339, 104, 1, 0,
          0, 0, 340, 342, 3, 113, 56, 0, 341, 340, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 341, 1, 0, 0,
          0, 343, 344, 1, 0, 0, 0, 344, 106, 1, 0, 0, 0, 345, 346, 7, 5, 0, 0, 346, 108, 1, 0, 0, 0, 347,
          348, 7, 6, 0, 0, 348, 110, 1, 0, 0, 0, 349, 350, 7, 7, 0, 0, 350, 112, 1, 0, 0, 0, 351, 352, 7,
          8, 0, 0, 352, 114, 1, 0, 0, 0, 353, 358, 5, 34, 0, 0, 354, 357, 3, 117, 58, 0, 355, 357, 8, 9,
          0, 0, 356, 354, 1, 0, 0, 0, 356, 355, 1, 0, 0, 0, 357, 360, 1, 0, 0, 0, 358, 356, 1, 0, 0, 0, 358,
          359, 1, 0, 0, 0, 359, 361, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 361, 362, 5, 34, 0, 0, 362, 116,
          1, 0, 0, 0, 363, 365, 5, 92, 0, 0, 364, 366, 7, 10, 0, 0, 365, 364, 1, 0, 0, 0, 366, 118, 1, 0,
          0, 0, 367, 368, 5, 47, 0, 0, 368, 369, 5, 47, 0, 0, 369, 373, 1, 0, 0, 0, 370, 372, 8, 11, 0,
          0, 371, 370, 1, 0, 0, 0, 372, 375, 1, 0, 0, 0, 373, 371, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374,
          376, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 376, 377, 6, 59, 0, 0, 377, 120, 1, 0, 0, 0, 378, 379,
          5, 47, 0, 0, 379, 380, 5, 42, 0, 0, 380, 384, 1, 0, 0, 0, 381, 383, 9, 0, 0, 0, 382, 381, 1, 0,
          0, 0, 383, 386, 1, 0, 0, 0, 384, 385, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 387, 1, 0, 0, 0, 386,
          384, 1, 0, 0, 0, 387, 388, 5, 42, 0, 0, 388, 389, 5, 47, 0, 0, 389, 390, 1, 0, 0, 0, 390, 391,
          6, 60, 0, 0, 391, 122, 1, 0, 0, 0, 392, 394, 7, 12, 0, 0, 393, 392, 1, 0, 0, 0, 394, 395, 1, 0,
          0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 397, 1, 0, 0, 0, 397, 398, 6, 61, 0, 0,
          398, 124, 1, 0, 0, 0, 14, 0, 295, 305, 313, 316, 323, 333, 343, 356, 358, 365, 373, 384,
          395, 1, 6, 0, 0};
  staticData->serializedATN = antlr4::atn::SerializedATNView(serializedATNSegment, sizeof(serializedATNSegment) / sizeof(serializedATNSegment[0]));

  antlr4::atn::ATNDeserializer deserializer;
  staticData->atn = deserializer.deserialize(staticData->serializedATN);

  const size_t count = staticData->atn->getNumberOfDecisions();
  staticData->decisionToDFA.reserve(count);
  for (size_t i = 0; i < count; i++) { 
    staticData->decisionToDFA.emplace_back(staticData->atn->getDecisionState(i), i);
  }
  riddlelexerLexerStaticData = staticData.release();
}

}

RiddleLexer::RiddleLexer(CharStream *input) : Lexer(input) {
  RiddleLexer::initialize();
  _interpreter = new atn::LexerATNSimulator(this, *riddlelexerLexerStaticData->atn, riddlelexerLexerStaticData->decisionToDFA, riddlelexerLexerStaticData->sharedContextCache);
}

RiddleLexer::~RiddleLexer() {
  delete _interpreter;
}

std::string RiddleLexer::getGrammarFileName() const {
  return "RiddleLexer.g4";
}

const std::vector<std::string>& RiddleLexer::getRuleNames() const {
  return riddlelexerLexerStaticData->ruleNames;
}

const std::vector<std::string>& RiddleLexer::getChannelNames() const {
  return riddlelexerLexerStaticData->channelNames;
}

const std::vector<std::string>& RiddleLexer::getModeNames() const {
  return riddlelexerLexerStaticData->modeNames;
}

const dfa::Vocabulary& RiddleLexer::getVocabulary() const {
  return riddlelexerLexerStaticData->vocabulary;
}

antlr4::atn::SerializedATNView RiddleLexer::getSerializedATN() const {
  return riddlelexerLexerStaticData->serializedATN;
}

const atn::ATN& RiddleLexer::getATN() const {
  return *riddlelexerLexerStaticData->atn;
}




void RiddleLexer::initialize() {
#if ANTLR4_USE_THREAD_LOCAL_CACHE
  riddlelexerLexerInitialize();
#else
  ::antlr4::internal::call_once(riddlelexerLexerOnceFlag, riddlelexerLexerInitialize);
#endif
}
