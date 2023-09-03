/** @format */

import { useEffect, useState } from "react";

// Custom hook to manage beerData state and update it with new data
const useBeerData = (initialData, isWebSocketReady, lastJsonMessage) => {
  const [beerData, setBeerData] = useState([]);

  // Function to filter and update beerData
  const updateBeerData = (newData) => {
    setBeerData((prevData) => [...prevData, ...newData]);
  };

  // Update beerData with initial data from the API
  useEffect(() => {
    if (initialData) {
      const filteredData = initialData.beverages.filter((beverage) => {
        return beverage.type.toString() === "BEVERAGE_TYPE_BEER";
      });
      setBeerData(filteredData);
    }
  }, [initialData]);

  // Listen to WebSocket messages and update beerData
  useEffect(() => {
    if (isWebSocketReady && lastJsonMessage) {
      // Assuming the WebSocket message contains new beer data
      const newBeerData = lastJsonMessage.data.filter((beverage) => {
        return beverage.type.toString() === "BEVERAGE_TYPE_BEER";
      });

      if (newBeerData.length > 0) {
        updateBeerData(newBeerData);
      }
    }
  }, [isWebSocketReady, lastJsonMessage]);

  return beerData;
};

export default useBeerData;
