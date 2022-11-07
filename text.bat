rd /s /q tmp
md tmp\blocks
md tmp\wallets
md tmp\ref_list

main.exe createwallet -refname A
main.exe createwallet -refname B
main.exe createwallet -refname C
main.exe walletslist
main.exe createblockchain -refname A
main.exe blockchaininfo
main.exe balance -refname A
main.exe sendbyrefname -from A -to B -amount 100
main.exe balance -refname B
main.exe mine
main.exe blockchaininfo
main.exe balance -refname A
main.exe balance -refname B
main.exe sendbyrefname -from A -to C -amount 100
main.exe sendbyrefname -from B -to C -amount 30
main.exe mine
main.exe blockchaininfo
main.exe balance -refname A
main.exe balance -refname B
main.exe balance -refname C
