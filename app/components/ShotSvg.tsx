/** @format */

import { View } from "react-native";
import Svg, { Path } from "react-native-svg";

const ShotSvg = () => (
  <View>
    <Svg
      width={24}
      height={24}
      fill="#000000"
      viewBox="0 0 16 16"
      xmlns="http://www.w3.org/2000/svg"
      className="inline-block w-10 h-10 mx-auto"
    >
      <Path d="m11.37 9.85.87-6.4a1.22 1.22 0 0 0-.3-1A1.24 1.24 0 0 0 11 2H2.85a1.24 1.24 0 0 0-1.23 1.41l1.15 8.48A1.27 1.27 0 0 0 4 13h4.84a4.41 4.41 0 0 0 2.66 1h.14a3.9 3.9 0 0 0 2.76-1.12zM11 3.25l-.15 1.09H3l-.15-1.09zM7.23 9.72a4.28 4.28 0 0 0 .57 2H4l-.83-6.13h7.47l-.42 3.11-1.87-1.87a3.83 3.83 0 0 0-1.12 2.89zm4.31 3a3.18 3.18 0 0 1-2.12-.94 3.21 3.21 0 0 1-.95-2.13 2.71 2.71 0 0 1 .11-.86l1.91 1.92 1.91 1.91a2.73 2.73 0 0 1-.76.1z" />
    </Svg>
  </View>
);

export default ShotSvg;
