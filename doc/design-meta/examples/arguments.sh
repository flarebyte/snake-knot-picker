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

# Boolean flags
wash start bedding --extra-rinse

# String Slists/Slices
wash start whites --add bleach --add softener
wash start whites --add bleach,softener,scent-beads --add scent-booster
# KO wash start whites --add=bleach

# comma-separated values and multiple flag instances
wash start --options delicate,extra-rinse --options pre-wash
