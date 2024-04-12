import { useEffect, useState } from "react";
import { Car } from "../types/models";
import { getCar } from "../services/car";
import { useParams } from "react-router";

const CarPage = () => {
  const params = useParams();
  console.debug({ params });
  const { id } = params;
  const [car, setCar] = useState<Car | undefined>();
  const [error, setError] = useState("");

  useEffect(() => {
    getCar(Number(id))
      .then((data) => setCar(data))
      .catch((err) => {
        setError("Error trying to fetch data");
        console.log(err);
      });
  }, [id]);

  if (error) return <p>{error}</p>;

  return <p>{car?.make}</p>;
};

export default CarPage;
