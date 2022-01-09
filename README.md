## Riakg Description

Is a **development oriented** command line tool that enables navigation on buckets, keys and values of a Riak KV data store using it's [HTTP interface](ttps://docs.riak.com/riak/kv/2.2.3/developing/api/http/index.html) using the awesome golang [tview](https://github.com/rivo/tview) library.

#### IMPORTANT NOTE:
RiakG purpose is to help on the development phase given the fact that there are not GUI interfaces that allows read or modify values on a Riak cluster, for that it uses the HTTP API and will list **ALL** the keys on a given bucket  so is not suitable for production.


## Installation

#### Mac OS:
if you are homebrew user you can install Riakg using a tap:

```bash
brew tap Heavyblade/riakg
brew install riakg
```
### Linux / Windows / Mac

You can download the pre-build binaries availables on the [releases section](https://github.com/Heavyblade/Riakg/releases) 

## Usage
By using the riakg command you will get a GUI with three sections ( buckets | keys | values ) and you can navigate through them using the TAB and arrow keys.

![riakg](https://github.com/Heavyblade/Riakg/blob/main/assets/riakg_reference.png)

features:
- key deletion (by pressing Ctrl+D when a key is selected)
- getting a key value (by pressing Ctrl-Y the value content of the key will be copied to the clipboard)
- Updating a value for key (by pressiong Ctrl-V the content on the clipboard will replace the current content on the value section.
- Saving the value (by pressing Ctrl-S the current value of the value text section will be sent to replace the current value of the key on the bucket) 

## License (MIT)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE
OR OTHER DEALINGS IN THE SOFTWARE.
