Deno.serve(async (req) => {
  // const { name } = await req.json();
  const data = {
    message: `Hello ${123}!`,
  };

  return new Response(JSON.stringify(data), {
    headers: { "Content-Type": "application/json" },
  });
});
