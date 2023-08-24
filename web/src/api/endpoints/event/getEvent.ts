/** @format */
import ENV from "@/constants/env";
import {
  GetEventRequest,
  GetEventResponse,
} from "@/api/protos/v1/event/event_pb";

import { useQuery } from "@tanstack/react-query";

import { http } from "@/api/axios";

if (ENV.MOCKED) {
  await import("./getEvent.mock.ts");
}

//This function I might actually be able to
const GetEvent = async (req: GetEventRequest) => {
  const { data } = await http.get(`v1/event/get`);

  return data as GetEventResponse;
};

const useGetEvent = (req: GetEventRequest) => {
  const { data, isLoading, error } = useQuery({
    queryKey: ["useGetEvent"],
    queryFn: () => GetEvent(req),
  });

  if (error) {
    console.error(error);
  }

  return { data, isLoading, error };
};

export default useGetEvent;
