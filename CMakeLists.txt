cmake_minimum_required(VERSION 3.28)
project(Riddle_Language)

set(CMAKE_CXX_STANDARD 20)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_SCAN_FOR_MODULES ON)

message("Current operating system: ${CMAKE_SYSTEM_NAME}")
message("Current C++ compiler: ${CMAKE_CXX_COMPILER_ID}")
message("Current C compiler: ${CMAKE_C_COMPILER_ID}")
message("Current build type: ${CMAKE_BUILD_TYPE}")
message("Current generator: ${CMAKE_GENERATOR}")

# Enter the system type for better searching of the database.
if(WIN32)
    set(PLATFORM_DEFINES "-DWIN32")
else(LINUX)
    set(PLATFORM_DEFINES "-DLINUX")
endif()

find_path(ANTLR4_INCLUDE_DIR antlr4-runtime/antlr4-runtime.h
        PATHS ${CMAKE_PREFIX_PATH} NO_DEFAULT_PATH)

find_package(LLVM REQUIRED CONFIG)
find_package(antlr4-runtime CONFIG REQUIRED)
find_path(TERMCOLOR_INCLUDE_DIRS "termcolor/termcolor.hpp")

message(STATUS "Found LLVM ${LLVM_PACKAGE_VERSION}")
message(STATUS "Using LLVMConfig.cmake in: ${LLVM_DIR}")
message(STATUS "Found Antlr4 include ${ANTLR4_INCLUDE_DIR}")

include_directories(${TERMCOLOR_INCLUDE_DIRS})
include_directories(src)
include_directories(${ANTLR4_INCLUDE_DIR})
include_directories(${LLVM_INCLUDE_DIRS})
include_directories(parser)

file(GLOB_RECURSE AST_FILES "parser/*.cpp" "parser/*.h")
file(GLOB_RECURSE SOURCE_FILES "src/*.cpp" "src/*.h")
file(GLOB_RECURSE MODULE_FILES "src/*.ixx")

add_library(Riddle_Module ${SOURCE_FILES})

target_sources(Riddle_Module PUBLIC FILE_SET CXX_MODULES FILES ${MODULE_FILES})

add_executable(Riddle_Language ${SOURCE_FILES} ${AST_FILES})

target_compile_definitions(Riddle_Language PRIVATE ${PLATFORM_DEFINES})

target_link_libraries(Riddle_Language LLVMCore LLVMLinker)
target_link_libraries(Riddle_Language Riddle_Module)
target_link_libraries(Riddle_Language antlr4_shared)
