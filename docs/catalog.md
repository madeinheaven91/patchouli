# Catalog

Adding a book:
POST catalog/v1/book
    {
        title: string,
        author: string,
        description: string,
        format: string,
        category_id: int,
        document: string
    }
