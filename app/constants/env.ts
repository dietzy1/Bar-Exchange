/** @format */
import Config from "react-native-config";

interface env {
  API_URL: string;
  MOCKED: boolean;
}

//Im lazy right now we need to read in from the .env file later
const ENV: env = {
  API_URL: "localhost:8000",
  MOCKED: true,
};

export default ENV;

/* Config.API_URL; // 'https://myapi.com'
Config.GOOGLE_MAPS_API_KEY; // 'abcdefgh' */
