# Utility Function to Check 2 Versions

#  Given 2 versions, version1 and version2, the compareVersions function will iterate through the a given version and return a 1 if version1 is "newer" than version2, a -1 if version1 is "older" than version2, or otherwise return a 0.

# "Newer" is defined as numbering that is more than the compared version. For example, version 1.5 is newer than version 1.0, and version 0.5, as its first level revision is greater than 0. 

# "Older" is defined as numbering that is less than the compared version. For example, version 12.0.9.9 is older than version 13.0.0.1 as the first level version of 13.0.0.1 (13) is greater than the first level version of 12.0.9.9 (12)

### Assumptions:
- Strings only contain digits and the "dot" (.) character
- Version values are below size 32-bit integers. Versions usually don't go this high.

### Complexity is in O(n) time and O(n) space where n is the length of the split version strings

### Future Changes
- Support integers above 32 bit by using parseInt
- Could split the string in half by possibly decrease the space, but will still be worst case O(n) space


