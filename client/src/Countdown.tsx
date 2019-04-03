import React, { useState, useEffect } from "react";
import axios from "axios";
import { useAsyncEffect } from "use-async-effect";

function Countdown() {
  const [lastPizza, setLastPizza] = useState();
  const [initialized, setInitialized] = useState(false);
  const [hours, setHours] = useState(0);
  const [minutes, setMinutes] = useState(0);
  const [seconds, setSeconds] = useState(0);
  const [days, setDays] = useState(0);

  useAsyncEffect(async () => {
    if (!initialized) {
      setInterval(function() {
        calculateTimeLeft();
      }, 1000);
      const result = await axios("api/time");
      setLastPizza(new Date(result.data.pizzaTime));
      setInitialized(true);
    }
  });
  const calculateTimeLeft = () => {
    if (lastPizza) {
      var myTimeSince = new Date().getTime() - lastPizza.getTime();
      setHours(differenceToHours(myTimeSince));
      setDays(differenceToDays(myTimeSince));
      setMinutes(differenceToMinutes(myTimeSince));
      setSeconds(differenceToSeconds(myTimeSince));
    }
  };

  const differenceToDays = (difference: number) => {
    return Math.floor(difference / (1000 * 60 * 60 * 24));
  };

  const differenceToHours = (difference: number) => {
    return Math.floor((difference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
  };
  const differenceToMinutes = (difference: number) => {
    return Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
  };
  const differenceToSeconds = (difference: number) => {
    return Math.floor((difference % (1000 * 60)) / 1000);
  };

  const updateLastPizza = async () => {
    const result = await axios("api/create");
    window.location.reload();
  };

  return (
    <section>
      <h2>Time Since Last Pizza (TSLP)</h2>
      <div>
        <p>Days: {days}</p>
        <p>Hours: {hours}</p>
        <p>Minutes: {minutes}</p>
        <p>Seconds: {seconds}</p>
      </div>
      <button onClick={updateLastPizza}>I just had pizza!! Click here</button>
    </section>
  );
}
export default Countdown;
