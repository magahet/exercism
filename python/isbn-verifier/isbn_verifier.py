def verify(isbn):
    # Validate string length
    if len(isbn.replace('-', '')) != 10:
        return False
        
    # Cast digits from string
    digits = [int(c) for c in isbn if c.isdigit()]
    if len(digits) == 9 and isbn.endswith('X'):
        digits.append(10)
    
    # Validate number of digits
    if len(digits) != 10:
        return False
        
    # Calculate checksum
    chksum = sum([a * b for a, b in zip(digits, xrange(10, 0, -1))])
    
    # Return checksum evaluation
    return chksum % 11 == 0