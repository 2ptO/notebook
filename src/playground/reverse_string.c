#include <stdio.h>
#include <string.h>

int reverse(char *word);

int main(void) {
    char input[128];
    printf("Enter your string to reverse:");
    scanf("%s", input);
    reverse(input);
    printf("Reversed string: %s\n", input);
    return 1;
}

/*
 * Given a word, reverse its characters.
 * abcd -> dcba
 */
int reverse(char *word)
{
    char *start = word;
    char *end = start;

    if (word == NULL) {
        return -1;
    }

    //find the end of the string.
    while (*end != '\0') {
        end++;
    }

    //go back one position in the end pointer to skip the NULL character.
    end--;

    //swap start and beginning
    while (start < end) {
        char tmp = *start;
        *start = *end;
        *end = tmp;
        ++start;
        --end;
    }

    return 1;
}
