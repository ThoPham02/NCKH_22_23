import { useEffect, useState } from 'react';

const NumberCounter = ({ target, duration, className }) => {
  const [count, setCount] = useState(0);

  useEffect(() => {
    let frameCount = 0;
    const framesPerSecond = 60;
    const increment = target / (duration * framesPerSecond);

    const timer = setInterval(() => {
      frameCount++;
      setCount(Math.round(increment * frameCount));

      if (frameCount === duration * framesPerSecond) {
        clearInterval(timer);
      }
    }, 1000 / framesPerSecond);

    return () => {
      clearInterval(timer);
    };
  }, [target, duration]);

  return <div className={className}>{count}</div>;
};

export default NumberCounter;