package main 

import piscine 


void my_print_comb(void)
{
    int i;
    int j;
    int k;
    for(i = '0'; i <= '9'; ++i)
    {
        for(j = '0'; j <= '9'; ++j)
        {
            for(k = '0'; k <= '9'; ++k) // k, not j, is tested against '9'
            {
                my_putchar(i);
                my_putchar(j);
                my_putchar(k);
                my_putchar(',');
                my_putchar(' ');
            }
        }
    }
}
