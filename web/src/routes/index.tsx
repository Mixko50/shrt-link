import { createEffect, createSignal } from "solid-js";
import {
  Dialog,
  DialogPanel,
  DialogTitle,
  Transition,
  TransitionChild,
  DialogOverlay,
} from "solid-headless";
import { BasedResponse, ErrorResponse, ShrtResponse } from "~/types/response";
import { RetrieveRequest, ShrtRequest } from "~/types/request";

const Home = () => {
  const [slug, setSlug] = createSignal<string>("");
  const [url, setUrl] = createSignal<string>("");

  const [response, setResponse] =
    createSignal<BasedResponse<ErrorResponse | ShrtResponse>>();

  const [isOpen, setIsOpen] = createSignal(false);

  const closeModal = () => {
    setIsOpen(false);
  };

  const openModal = () => {
    setIsOpen(true);
  };

  const createShortUrl = async () => {
    const req: ShrtRequest = {
      full_url: url(),
      slug: slug(),
    };
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(req),
    };
    fetch(import.meta.env.VITE_BASE_URL + "/api/create", requestOptions)
      .then((response) => response.json())
      .then((data) => setResponse(data));
  };

  const retrieveOriginalUrl = () => {
    const req: RetrieveRequest = {
      shrt_url: url(),
    };
    const requestOptions = {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(req),
    };
    fetch(import.meta.env.VITE_BASE_URL + "/api/retrive", requestOptions)
      .then((response) => response.json())
      .then((data) => setResponse(data));
  };

  createEffect(() => {
    setTimeout(() => {
      openModal();
    }, 1000);
  });

  return (
    <div class="bg-white">
      <div class="relative isolate px-6 pt-14 lg:px-8">
        <div class="absolute inset-x-0 -top-40 -z-10 transform-gpu overflow-hidden blur-3xl sm:-top-80">
          <svg
            class="relative left-[calc(50%-11rem)] -z-10 h-[21.1875rem] max-w-none -translate-x-1/2 rotate-[30deg] sm:left-[calc(50%-30rem)] sm:h-[42.375rem]"
            viewBox="0 0 1155 678"
          >
            <path
              fill="url(#45de2b6b-92d5-4d68-a6a0-9b9b2abad533)"
              fill-opacity=".3"
              d="M317.219 518.975L203.852 678 0 438.341l317.219 80.634 204.172-286.402c1.307 132.337 45.083 346.658 209.733 145.248C936.936 126.058 882.053-94.234 1031.02 41.331c119.18 108.451 130.68 295.337 121.53 375.223L855 299l21.173 362.054-558.954-142.079z"
            />
            <defs>
              <linearGradient
                id="45de2b6b-92d5-4d68-a6a0-9b9b2abad533"
                x1="1155.49"
                x2="-78.208"
                y1=".177"
                y2="474.645"
                gradientUnits="userSpaceOnUse"
              >
                <stop stop-color="#9089FC" />
                <stop offset="1" stop-color="#FF80B5" />
              </linearGradient>
            </defs>
          </svg>
        </div>
        {/*  */}
        <div class="mx-auto max-w-2xl py-32 sm:py-48 lg:py-56">
          <div class="text-center">
            <h1 class="text-4xl font-bold tracking-tight text-gray-900 sm:text-6xl">
              Shrt URL
            </h1>
            <p class="mt-6 text-lg leading-8 text-gray-600">URL shortener</p>

            <div class="mt-5">
              <div class="mt-5 md:col-span-2 md:mt-0">
                <div class="overflow-hidden shadow sm:rounded-md">
                  <div class="bg-white px-4 py-5 sm:p-6">
                    <div class="grid grid-cols-6 gap-6">
                      <div class="col-span-6 sm:col-span-4">
                        <label class="block text-sm font-medium leading-6 text-gray-900">
                          URL
                        </label>
                        <input
                          type="text"
                          class="mt-2 block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                          onInput={(e) =>
                            setUrl((e.target as HTMLInputElement).value)
                          }
                        />
                      </div>

                      <div class="col-span-6 sm:col-span-2">
                        <label class="block text-sm font-medium leading-6 text-gray-900">
                          Slug
                        </label>
                        <input
                          type="text"
                          class="mt-2 block w-full rounded-md border-0 py-1.5 px-2 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
                          onInput={(e) =>
                            setSlug((e.target as HTMLInputElement).value)
                          }
                        />
                      </div>
                    </div>
                  </div>
                  <div class="pb-5 space-x-2">
                    <button
                      type="submit"
                      class="h-10 rounded-md bg-[#E96479] py-2 px-3 text-sm font-semibold text-white shadow-sm hover:bg-[#C3ACD0] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-500"
                    >
                      Get shorten url
                    </button>
                    <button
                      type="submit"
                      class="h-10 rounded-md bg-[#7286D3] py-2 px-3 text-sm font-semibold text-white shadow-sm hover:bg-[#C3ACD0] focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-[indigo-500]"
                    >
                      Retrieve original url
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
          {/* / */}
        </div>
        <div class="absolute inset-x-0 top-[calc(100%-13rem)] -z-10 transform-gpu overflow-hidden blur-3xl sm:top-[calc(100%-30rem)]">
          <svg
            class="relative left-[calc(50%+3rem)] h-[21.1875rem] max-w-none -translate-x-1/2 sm:left-[calc(50%+36rem)] sm:h-[42.375rem]"
            viewBox="0 0 1155 678"
          >
            <path
              fill="url(#ecb5b0c9-546c-4772-8c71-4d3f06d544bc)"
              fill-opacity=".3"
              d="M317.219 518.975L203.852 678 0 438.341l317.219 80.634 204.172-286.402c1.307 132.337 45.083 346.658 209.733 145.248C936.936 126.058 882.053-94.234 1031.02 41.331c119.18 108.451 130.68 295.337 121.53 375.223L855 299l21.173 362.054-558.954-142.079z"
            />
            <defs>
              <linearGradient
                id="ecb5b0c9-546c-4772-8c71-4d3f06d544bc"
                x1="1155.49"
                x2="-78.208"
                y1=".177"
                y2="474.645"
                gradientUnits="userSpaceOnUse"
              >
                <stop stop-color="#9089FC" />
                <stop offset="1" stop-color="#FF80B5" />
              </linearGradient>
            </defs>
          </svg>
        </div>
      </div>

      {/* Dialog */}
      <Transition appear show={isOpen()}>
        <Dialog
          isOpen
          class="fixed inset-0 z-10 overflow-y-auto"
          onClose={closeModal}
        >
          <div class="min-h-screen px-4 flex items-center justify-center">
            <TransitionChild
              enter="ease-out duration-300"
              enterFrom="opacity-0"
              enterTo="opacity-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100"
              leaveTo="opacity-0"
            >
              <DialogOverlay class="fixed inset-0 bg-gray-900 bg-opacity-50" />
            </TransitionChild>

            {/* This element is to trick the browser into centering the modal contents. */}
            <span class="inline-block h-screen align-middle" aria-hidden="true">
              &#8203;
            </span>
            <TransitionChild
              enter="ease-out duration-300"
              enterFrom="opacity-0 scale-95"
              enterTo="opacity-100 scale-100"
              leave="ease-in duration-200"
              leaveFrom="opacity-100 scale-100"
              leaveTo="opacity-0 scale-95"
            >
              <DialogPanel class="inline-block w-full max-w-md p-6 my-8 overflow-hidden text-left align-middle transition-all transform bg-gray-50 shadow-xl rounded-xl dark:border dark:border-gray-50">
                <DialogTitle
                  as="h3"
                  class="text-xl font-medium leading-6 text-gray-900"
                >
                  {response()?.success ? "Success!" : "Failed"}
                </DialogTitle>
                <div class="mt-2 flex space-x-2">
                  <p class="text-lg font-medium">Your shrt link:</p>
                  <p class="text-lg">http://sdsdsd</p>
                </div>

                <div class="mt-4 flex justify-center space-x-2">
                  <button
                    type="button"
                    class="w-full justify-center h-10 rounded-md bg-[#C780FA] px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto "
                    onClick={closeModal}
                  >
                    Copy
                  </button>
                  <button
                    type="button"
                    class="w-full justify-center h-10 rounded-md bg-[#FFB26B] px-3 py-2 text-sm font-semibold text-white shadow-sm hover:bg-red-500 sm:ml-3 sm:w-auto "
                    onClick={closeModal}
                  >
                    Open url
                  </button>
                </div>
              </DialogPanel>
            </TransitionChild>
          </div>
        </Dialog>
      </Transition>
    </div>
  );
};

export default Home;
