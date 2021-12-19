package gen

var EditorConfigData = []byte(`
root = true
[*]
end_of_line = lf
insert_final_newline = true
indent_style = space
charset = utf-8
trim_trailing_whitespace = true
indent_size = 2
##################################
####### Made by Moldy CLI ########
##################################
`)

var GitIgnoreData = []byte(`
# Ignoring the .env files

*.env

# Editor ignore folders
*.idea/
*.vscode/
*.vim/

##################################
####### Made by Moldy CLI ########
##################################
`)

var ReadmeData = []byte(`
# Example package

Example content and example description :p

# Credits

**Author:** Jhon Doe
**Version:** v1.0.0
**Url:** [here](https://example.com)

> Made with :heart: by Moldy CLI [here](https://github.com/Moldy-Community/moldy)
`)

var LicenseData = []byte(`

Copyright © 2021 Example Author
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
`)
