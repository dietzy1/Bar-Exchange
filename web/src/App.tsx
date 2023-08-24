/** @format */

import ReactQueryClientProvider from "./api/queryClient";
import Homepage from "./views/Homepage";

function App() {
  return (
    <ReactQueryClientProvider>
      <Homepage />
    </ReactQueryClientProvider>
  );
}

export default App;
