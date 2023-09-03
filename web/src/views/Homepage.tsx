/** @format */

import Header from "../components/Header";
import Footer from "../components/Footer";
import { Datatable } from "../components/Datatable";

import { columns } from "../components/Columns";

import CocktailSvg from "../components/CocktailSvg";
import BeerSvg from "../components/BeerSvg";
import ShotSvg from "../components/ShotSvg";
import useWebsocket from "@/api/websocket/useWebsocket";
import useGetBeverages from "@/api/endpoints/beverage/getBeverages";
import Spinner from "@/components/Spinner";

const Homepage = () => {
  // const [data, setData] = React.useState<Beverage[]>([]);

  const { data, isLoading, error } = useGetBeverages();
  if (error) {
    console.log(error);
  }

  const { lastJsonMessage, isWebSocketReady } = useWebsocket();

  const beerData = data?.beverages.filter((beverage) => {
    //Map beverageType to the correct enum number
    return beverage.type.toString() === "BEVERAGE_TYPE_BEER";
  });

  const cocktailData = data?.beverages.filter((beverage) => {
    //Map beverageType to the correct enum number
    return beverage.type.toString() === "BEVERAGE_TYPE_COCKTAIL";
  });

  const shotsData = data?.beverages.filter((beverage) => {
    //Map beverageType to the correct enum number
    return beverage.type.toString() === "BEVERAGE_TYPE_SHOTS";
  });

  if (isWebSocketReady) {
    console.log(lastJsonMessage);
  }

  return (
    <div className="w-screen h-screen dark:bg-background flex flex-col">
      <Header />

      <div className="container pt-8 flex-row flex justify-center ">
        {isLoading ? (
          <Spinner />
        ) : (
          <div className="rounded-3xl p-6 flex">
            <div className="mx-5 flex flex-col">
              <div className="flex flex-row self-center">
                <BeerSvg />
                <h1 className="text-4xl self-center mb-2">Øl</h1>
              </div>
              <Datatable columns={columns} data={beerData!} />
            </div>
            <div className="mx-5 flex flex-col ">
              <div className="flex flex-row self-center">
                <CocktailSvg />
                <h1 className="text-4xl self-center mb-2">Drinks</h1>
              </div>
              <Datatable columns={columns} data={cocktailData!} />
            </div>
            <div className="mx-5 flex flex-col mb-2">
              <div className="flex flex-row self-center">
                <ShotSvg />
                <h1 className="text-4xl self-center mb-2">Shots</h1>
              </div>
              <Datatable columns={columns} data={shotsData!} />
            </div>
          </div>
        )}
      </div>

      <div className="flex flex-row justify-center">
        <div className="flex flex-col"></div>
      </div>

      <Footer />
    </div>
  );
};

export default Homepage;
