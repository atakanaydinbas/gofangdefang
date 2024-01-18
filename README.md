# GoFangDefang

GoFangDefang is a simple Go library designed to handle IOC manipulation by converting special characters within a IOC from "fang" format to "defang" format and vice versa. It is a process to prevent clicking on potentially malicious clickable IOC's such as some url, ip, domain or subdomain.

## What are Fang and Defang?

- **Fang Format**: Refers to an IOC with special characters such as periods (`.`, `https://`, `@`) in their original form.
- **Defang Format**: Refers to the same IOC or text, but the special characters are modified, preventing them from being executed or interpreted by systems. For example, replacing `.` with `[.]`, replacing `@` with `[AT]`, or replacing `https://` with `hXXps://`

## Features

- **Fang to Defang**: Convert a IOC from its original "fang" format to a more secure "defang" format.
- **Defang to Fang**: Convert a IOC from its "defang" format back to its original "fang" format.
- **Ease of Integration**: Easily integrate this tool into your security scripts or applications for IOC manipulation.

## Usage

To use GoFangDefang in your Go projects, import it directly from GitHub:
- ``` go get github.com/atakanaydinbas/gofangdefang```

```go
package main

import (
    "fmt"
    "github.com/atakanaydinbas/gofangdefang"
)

func main() {
    // Fanging a URL
    fangedURL := gofangdefang.FangAll("hXXps://example[.]com")
    // Defanging an email address
    defangedEmail := gofangdefang.DefangAll("foo@example.com")

    // Printing the results
    fmt.Println(fangedURL)      // --> https://example.com
    fmt.Println(defangedEmail)  // --> foo[AT]example[.]com

    // Example usage for file operations
    filePath := "path/to/your/file.txt"

    // Fanging IOCs in a file
    // FangFile is used to process IOCs in a file.
    // Parameters:
    //   - filepath string: Path to the file containing IOCs.
    //   - willbesavedfile bool: Specifies if the processed data will be saved to a new file or not.
    //   - newfilename ...string: Vararg, optionally used to specify the new filename.
    // Returned Values:
    //   - string: If willbesavedfile parameter is false, it will be text file of given file. If willbesavedfile parameter is true, it will be ""
    //   - error: Error message in case of an error.
    fangedFile, err := gofangdefang.FangFile(filePath, true, "fanged_file.txt")
    if err != nil {
        fmt.Println("FangFile Error:", err)
        return
    }

    // Defanging IOCs in a file
    // DefangFile is used to process IOCs in a file.
    // Parameters:
    //   - filepath string: Path to the file containing IOCs.
    //   - willbesavedfile bool: Specifies if the processed data will be saved to a new file or not.
    //   - newfilename ...string: Vararg, optionally used to specify the new filename.
    // Returned Values:
    //   - string: If willbesavedfile parameter is false, it will be text file of given file. If willbesavedfile parameter is true, it will be ""
    //   - error: Error message in case of an error.
    defangedFile, err := gofangdefang.DefangFile(filePath, true, "defanged_file.txt")
    if err != nil {
        fmt.Println("DefangFile Error:", err)
        return
    }
}
```

## Contribution
Contributions to GoFangDefang are welcome! If you'd like to contribute, please fork the repository and create a pull request. Ensure your code follows the project's style and conventions.
