/** @format */

import React from "react";

import Header from "./components/Header";
import Footer from "./components/Footer";
import { Datatable } from "./components/Datatable";

import { useEffect } from "react";
import { Beverage, columns } from "./components/Columns";

import CocktailSvg from "./components/CocktailSvg";
import BeerSvg from "./components/BeerSvg";
import ShotSvg from "./components/ShotSvg";

async function getData(): Promise<Beverage[]> {
  // Fetch data from your API here.
  return [
    {
      id: "728ed52f",
      item: "Blå vand",
      kind: "Drinks",
      amount: 76,
      status: "increasing",
      PercentageChange: 0.5,
    },
    {
      id: "489e1d42",
      item: "Ceres Top",
      kind: "Øl",
      amount: 125,
      status: "decreasing",
      PercentageChange: -0.1,
    },
    {
      id: "f3b1c2a3",
      item: "Royal Export",
      kind: "Øl",
      amount: 125,
      status: "decreasing",
      PercentageChange: 0.2,
    },
    {
      id: "f3b1c2a3",
      item: "Kongen af Danmark",
      kind: "Drinks",
      amount: 125,
      status: "decreasing",
      PercentageChange: 0.3,
    },
    {
      id: "f3b1c2a3",
      item: "Arnbitter",
      kind: "Shots",
      amount: 125,
      status: "decreasing",
      PercentageChange: 0.4,
    },
  ];
}

function App() {
  const [data, setData] = React.useState<Beverage[]>([]);

  useEffect(() => {
    (async function () {
      const data = await getData();
      setData(data);
    })();
  }, []);

  return (
    <div className="w-screen h-screen dark:bg-background flex flex-col">
      <Header />

      <div className="container pt-8 flex-row flex justify-center ">
        <div className="rounded-3xl p-6 flex">
          <div className="mx-5 flex flex-col">
            <div className="flex flex-row self-center">
              <BeerSvg />
              <h1 className="text-4xl self-center mb-2">Øl</h1>
            </div>
            <Datatable columns={columns} data={data} />
          </div>
          <div className="mx-5 flex flex-col ">
            <div className="flex flex-row self-center">
              <CocktailSvg />
              <h1 className="text-4xl self-center mb-2">Drinks</h1>
            </div>
            <Datatable columns={columns} data={data} />
          </div>
          <div className="mx-5 flex flex-col mb-2">
            <div className="flex flex-row self-center">
              <ShotSvg />
              <h1 className="text-4xl self-center mb-2">Shots</h1>
            </div>
            <Datatable columns={columns} data={data} />
          </div>
        </div>
      </div>

      <div className="flex flex-row justify-center">
        <div className="flex flex-col"></div>
      </div>

      <Footer />
    </div>
  );
}

export default App;
