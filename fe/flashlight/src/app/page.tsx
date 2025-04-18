"use client";
import { Pencil, Trash2, X } from "lucide-react";
import { useRef } from "react";

export default function Home() {
  const headers = ["ID", "Name", "Grade", "Actions"];

  const dialogCreateEditRef = useRef<HTMLDialogElement>(null);

  const onNewStudent = () => {
    dialogCreateEditRef.current?.showModal();
  };

  return (
    <div className="h-screen p-6">
      <dialog
        ref={dialogCreateEditRef}
        className="bg-white w-lg mx-auto mt-12 shadow-lg px-6 p-4 rounded-lg relative"
      >
        <button
          onClick={() => {
            dialogCreateEditRef.current?.close();
          }}
          className="absolute text-zinc-300 hover:text-zinc-700 duration-200 p-1.5 rounded cursor-pointer outline-none top-3 right-3 hover:bg-zinc-100"
        >
          <X className="w-4 h-4" />
        </button>
        <form>
          <h1 className="text-lg font-medium py-1">Create Student</h1>
          <label className="flex text-sm flex-col">
            Name
            <input
              required
              className="border peer invalid:border-red-400 duration-200 outline-none rounded-md border-gray-200 py-1.5 px-3 text-sm"
              type="text"
            />
            <p className="peer-invalid:visible text-xs text-red-400 invisible">
              This field is required
            </p>
          </label>

          <label className="flex mt-1 text-sm flex-col">
            Grade
            <input
              required
              className="border peer invalid:border-red-400 duration-200 outline-none rounded-md border-gray-200 py-1.5 px-3 text-sm"
              type="number"
              min={0}
            />
            <p className="peer-invalid:visible text-xs text-red-400 invisible">
              This field is required and needs to be bigger or equal to 0
            </p>
          </label>
          <div className="mt-4 flex gap-2">
            <div className="flex-1">
              <button
                type="button"
                onClick={() => dialogCreateEditRef.current?.close()}
                className="font-medium outline-none py-1.5 cursor-pointer text-gray-400 hover:text-gray-700 hover:bg-zinc-100 duration-200 px-3 rounded "
              >
                Cancel
              </button>
            </div>
            <div className="flex items-center gap-2">
              <button
                type="reset"
                className="font-medium outline-none py-1.5 cursor-pointer text-gray-400 hover:text-gray-700 hover:bg-zinc-100 duration-200 px-3 rounded "
              >
                Clear
              </button>
              <button
                onClick={onNewStudent}
                className="font-medium py-1.5 outline-none cursor-pointer hover:-translate-y-0.5 shadow-md hover:shadow-lg duration-200 px-3 rounded bg-amber-400"
              >
                Create
              </button>
            </div>
          </div>
        </form>
      </dialog>
      <header className="flex mt-3 gap-2 items-center">
        <div className="flex-1 flex flex-col gap-1">
          <h1 className="font-medium leading-none text-lg">Students</h1>
          <p>Assignment for the Flashlight Test</p>
        </div>
        <button
          onClick={onNewStudent}
          className="font-medium py-1.5 cursor-pointer outline-none hover:-translate-y-0.5 shadow-md hover:shadow-lg duration-200 px-3 rounded bg-amber-400"
        >
          New Student
        </button>
      </header>
      <table className="mt-7 w-full">
        <thead>
          <tr className="border-t border-zinc-300 border-b ">
            {headers.map((h, i) => (
              <th
                className="last-of-type:text-right font-medium py-2 last-of-type:px-3 first-of-type:px-3 first-of-type:text-left text-left [&:nth-of-type(2)]:w-2/4 "
                key={i}
              >
                {h}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          <tr className="even:bg-zinc-50">
            <td>1</td>
            <td>Brian Ito</td>
            <td>9</td>
            <td className="text-right py-2">
              <button className="p-2 rounded cursor-pointer mx-2 text-amber-400 hover:text-amber-600 duration-200 hover:bg-zinc-100">
                <Pencil className="w-4 h-4" />
              </button>
              <button className="p-2 rounded cursor-pointer text-red-300 hover:text-red-500 duration-200 hover:bg-zinc-100">
                <Trash2 className="w-4 h-4" />
              </button>
            </td>
          </tr>

          <tr className="even:bg-zinc-50">
            <td>1</td>
            <td>Brian Ito</td>
            <td>9</td>
            <td className="text-right py-2">
              <button className="p-2 rounded cursor-pointer mx-2 text-amber-400 hover:text-amber-600 duration-200 hover:bg-zinc-100">
                <Pencil className="w-4 h-4" />
              </button>
              <button className="p-2 rounded cursor-pointer text-red-300 hover:text-red-500 duration-200 hover:bg-zinc-100">
                <Trash2 className="w-4 h-4" />
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
