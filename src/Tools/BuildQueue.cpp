#include "BuildQueue.h"
#include "RiddleLexer.h"
#include "Visitors/GenVisitor.h"
#include "Visitors/PackageVisitor.h"
#include <iostream>

namespace Riddle {
    void BuildQueue::push(const Unit &unit) {
        libSource[unit.getPackName()].push_back(unit);
        for(auto i: unit.getImports()) {
            libGraph[i].push_back(unit.getPackName());
        }
    }
    void BuildQueue::start() {
        // todo 实现解析包相关的东西
        // 暂时还不写，先完成编译main
        if(libSource.count("main")) {
            GenVisitor visitor("main");
            visitor.visit(libSource["main"].front().parseTree);
        } else {
            std::cerr << R"(Not Found "main" package)" << std::endl;
            return;
        }
    }
    // 目前只做了main的解析
    void BuildQueue::parserFile(std::string filePath) {
        std::ifstream stream(filePath);
        const auto input = new antlr4::ANTLRInputStream(stream);
        const auto lexer = new RiddleLexer(input);
        const auto tokens = new antlr4::CommonTokenStream(lexer);
        auto *parser = new RiddleParser(tokens);
        antlr4::tree::ParseTree *p = parser->program();
        PackageVisitor visitor(filePath, p);
        push(visitor.unit);
        return;
    }
}// namespace Riddle