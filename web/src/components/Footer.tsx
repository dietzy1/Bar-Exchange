/** @format */

import { AiOutlineStock } from "react-icons/ai";

import Timer from "./Timer";

export default function Footer(): JSX.Element {
  return (
    <footer className="px-6 pb-4 w-screen mt-auto">
      <div className="text-1xl font-bold whitespace-nowrap text-white flex items-center justify-between pl-14 pr-14">
        <div className=" flex items-center text-black">
          <AiOutlineStock className="h-8 w-8 mr-2 text-black" />
          BØRSBAR
        </div>

        <div>
          <div className="flex flex-col items-center">
            <h1 className="text-2xl md:text-4xl font-bold mb-2 text-black">
              <span className="text-green-500">Børsen</span> lukker om
            </h1>
          </div>
          <div className="flex justify-center align-middle">
            <Timer />
          </div>
        </div>

        <div className=" flex items-center text-black invisible">
          <AiOutlineStock className="h-8 w-8 mr-2 text-black" />
          BØRSBAR
        </div>
      </div>

      <hr className="my-3 sm:mx-auto  " />
      <span className="block text-sm text-gray-400 sm:text-center ">
        © 2023 BØRSBAR™
      </span>
    </footer>
  );
}
