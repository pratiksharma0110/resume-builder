# Resume Builder CLI

A simple command-line interface (CLI) tool to generate a professional resume dynamically based on user input. The tool uses LaTeX for formatting and `pdflatex` to compile the `.tex` file into a PDF.

---

## Features

- Interactive CLI prompts to gather user information.
- Generates a well-formatted LaTeX `.tex` file for the resume.
- Compiles the `.tex` file into a PDF using `pdflatex`.
- Fully customizable template for personalizing the resume.

---

## Prerequisites

- **Go** installed on your system.
- **LaTeX** distribution installed (for `pdflatex`).
- Ensure `pdflatex` is available in your system path.

To check if `pdflatex` is installed, run:

```bash
pdflatex --version
