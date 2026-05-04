# Positional Arguments
wash start heavy-duty

# String Flags
wash start delicate --temp warm

# Integer flags
wash start normal --spin 1200
wash start normal --spin=1200
# OK ["--spin", "1200", "cottons"]
# OK ["--spin=1200", "cottons"]
# KO ["--spin 1200", "cottons"]

wash --add=bleach,softener,scent-beads
wash --add=bleach --add=softener

# Boolean flags
wash start bedding --extra-rinse

# String Slists/Slices
wash start whites --add bleach --add softener

# comma-separated values and multiple flag instances
wash start --options delicate,extra-rinse --options pre-wash

