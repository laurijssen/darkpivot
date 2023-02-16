#include<stdio.h>
#include<string.h>
#include<stdlib.h>
#include<unistd.h>
#include<sys/types.h>
#include<sys/wait.h>
#include<readline/readline.h>
#include<readline/history.h>

#define MAXCOM 1000 // max number of letters to be supported
#define MAXLIST 100 // max number of commands to be supported
  
#define clear() printf("\033[H\033[J")
  
void init_shell()
{
    clear();
    printf("\n\n******************"
        "************************");
    printf("\n\I love nGummibears!");
    printf("\n\n*******************"
        "***********************");
    printf("\n\nwelcome\n");
    sleep(1);
    clear();
}

int main(int args, char **argv)
{
    init_shell();

    return 0;
}
