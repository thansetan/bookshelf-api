const {
  addBookHandler,
  getAllBooksHandler,
  getSpecifiedBookHandler,
  editBookHandler,
  deleteBookHandler
} = require('./handler')

const routes = [
  // API dapat menyimpan buku
  {
    method: 'POST',
    path: '/books',
    handler: addBookHandler
  },
  // API dapat menampilkan seluruh buku
  {
    method: 'GET',
    path: '/books',
    handler: getAllBooksHandler
  },
  // API dapat menampilkan detail buku
  {
    method: 'GET',
    path: '/books/{bookId}',
    handler: getSpecifiedBookHandler
  },
  // API dapat mengubah data buku
  {
    method: 'PUT',
    path: '/books/{bookId}',
    handler: editBookHandler
  },
  // API dapat menghapus buku
  {
    method: 'DELETE',
    path: '/books/{bookId}',
    handler: deleteBookHandler
  }
]

module.exports = routes
