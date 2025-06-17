# Joinself Academy

Welcome to the Joinself Academy! This repository is the central hub for all educational materials related to the Joinself ecosystem.

Our mission is to make decentralized identity accessible to all developers by providing a clear, structured, and hands-on learning experience.

## Repository Structure

This repository is a monorepo containing the following key areas:

- **/docs**: The source content for our public documentation website, [academy.joinself.com](https://academy.joinself.com) (coming soon). The site is built with [MkDocs](https://www.mkdocs.org/) and the [Material for MkDocs](https://squidfunk.github.io/mkdocs-material/) theme.
- **/sdks**: Contains the source code for the developer-friendly client facades for various languages.
  - `go/`: Go client facade.
  - `java/`: (Placeholder) Java client facade.
  - `mobile/`: (Placeholder) Mobile (Swift/Kotlin) client facades.
- **/examples**: Houses practical code examples that correspond to the documentation. These are structured to provide a progressive learning path.
  - `go/`: Examples for the Go SDK and client facade.
  - `java/`: (Placeholder) Examples for the Java client facade.
  - `mobile/`: (Placeholder) Examples for the mobile client facades.

## Getting Started with the Documentation

To build and view the documentation locally:

1.  **Install dependencies:**
    ```bash
    python3 -m venv venv
    source venv/bin/activate
    pip install -r requirements.txt
    ```

2.  **Start the local server:**
    ```bash
    mkdocs serve
    ```

    The site will be available at `http://127.0.0.1:8000`.

## Contributing

We welcome contributions! Please see the `CONTRIBUTING.md` file (coming soon) for details on how to get started. 
