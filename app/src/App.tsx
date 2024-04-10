import { useEffect, useState } from "react";
import { getCars } from "./services/car";
import { Car } from "./types/models";

function App() {
  const [cars, setCars] = useState<Car[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    getCars()
      .then((data) => setCars(data))
      .catch((err) => {
        setError("Error trying to fetch data");
        console.log(err);
      });
  }, []);

  if (error) return <p>{error}</p>;

  return (
    <ul>
      {cars.map((car) => (
        <li key={car.ID}>{car.make}</li>
      ))}
    </ul>
  );
}

export default App;
