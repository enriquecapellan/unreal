import { useEffect, useState } from "react";

function App() {
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  const [books, setBooks] = useState<any[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/cars")
      .then((res) => res.json())
      .then((data) => setBooks(data));
  }, []);

  return (
    <ul>
      {books.map((book) => (
        <li key={book.ID}>{book.make}</li>
      ))}
    </ul>
  );
}

export default App;
