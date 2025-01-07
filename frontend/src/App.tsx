import { Alert, LoadingOverlay, Stack, Title } from "@mantine/core";
import { useGetUser } from "./API/useUser";

function App() {
  const { data, isLoading, isError, error } = useGetUser();
  return (
    <>
      {isError && <Alert>{error.message}</Alert>}
      <LoadingOverlay visible={isLoading} />
      <Stack>
        <Title>{data && data.user}</Title>
      </Stack>
    </>
  );
}

export default App;
