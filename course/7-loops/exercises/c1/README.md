PACKET SIZE
Textio needs to divide messages up into data packets. Normally, packet sizes are uniform and divided into bytes. An intern had the bright idea to customize packet sizes per message. Help textio test this terrible, terrible idea.

ASSIGNMENT
Implement the getPacketSize function to take a string message and return an integer. Find the largest packet size where the resulting number of packets is a composite number, ie: not 1 or prime. Use the provided isPrime function within getPacketSize. If no such packet size exists, return 0.

EXAMPLE
If the length of the message string is 64, its divisors are 64, 32, 16, 8, 4, 2 and 1. For packet sizes of 64 and 32, the number of packets would be 1 and 2 respectively, which are not composite. A packet size of 16, however, would result in 4 packets which is composite, so we return 16 as the greatest possible packet size.

getPacketSize("You shall not pass!!")
// returns 5

getPacketSize("May the odds be ever in your favor!!")
// returns 9
