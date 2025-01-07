import { useQuery } from "@tanstack/react-query";
import { axiosClient } from "./mainClient";

interface IGetUser {
  user: string;
}

export const useGetUser = () => {
  const getUsers = async (): Promise<IGetUser> => {
    const { data } = await axiosClient.get<IGetUser>("/user");
    return data;
  };

  return useQuery<IGetUser>({ queryKey: ["user"], queryFn: getUsers });
};
