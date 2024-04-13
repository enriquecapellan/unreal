import { useEffect, useState } from "react";
import { getCars } from "../services/car";
import { Car } from "../types/models";
import { Link } from "react-router-dom";

const CarsPage = () => {
  const [cars, setCars] = useState<Car[]>([]);
  const [error, setError] = useState("");

  useEffect(() => {
    getCars()
      .then((data) => setCars(data))
      .catch((err) => {
        setError("Error trying to fetch cars");
        console.log(err);
      });
  }, []);

  if (error) return <p>{error}</p>;

  return (
    <ul>
      {cars.map((car) => (
        <li key={car.ID}>
          <Link to={`/cars/${car.ID}`}>{car.make}</Link>
        </li>
      ))}
    </ul>
  );
};

export default CarsPage;
