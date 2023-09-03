/** @format */

import { FaSpinner } from "react-icons/fa";

const Spinner = () => {
  return (
    <div className="flex items-center justify-center">
      <FaSpinner className="animate-spin h-16 w-16 text-green-500 " />
    </div>
  );
};

export default Spinner;
