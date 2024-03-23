import { useEffect, useState } from "react";
import { getCars } from "./services/car";
import { Car } from "./types/models";

function App() {
  const [books, setBooks] = useState<Car[]>([]);

  useEffect(() => {
    getCars().then((data) => setBooks(data));
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
