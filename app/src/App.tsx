import { useEffect, useState } from "react";
import { getCars } from "./services/car";
import { Car } from "./types/models";

function App() {
  const [books, setBooks] = useState<Car[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    getCars()
      .then((data) => setBooks(data))
      .catch((err) => {
        setError("Error trying to fetch data");
        console.log(err);
      });
  }, []);

  if (error) return <p>{error}</p>;
  
  return (
    <ul>
      {books.map((book) => (
        <li key={book.ID}>{book.make}</li>
      ))}
    </ul>
  );
}

export default App;
