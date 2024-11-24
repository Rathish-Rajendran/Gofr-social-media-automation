import os
from langchain_community.vectorstores import Chroma
from langchain_ollama import OllamaEmbeddings
import dotenv
from git import Repo
from langchain_ollama import OllamaEmbeddings
from langchain_community.document_loaders import UnstructuredMarkdownLoader
from langchain_text_splitters import RecursiveCharacterTextSplitter

dotenv.load_dotenv()
current_dir = os.path.dirname(os.path.abspath(__file__))
persistent_directory = os.path.join(current_dir, "db", "chroma_db_llama32_claude")
current_dir = os.path.dirname(os.path.abspath(__file__))
local_path = f"{current_dir}/gofr_repo"

repo_url = "https://github.com/gofr-dev/gofr.git"
if not os.path.exists(local_path):
    Repo.clone_from(repo_url, local_path)
    print(f"Cloned {repo_url} into {local_path}")

embeddings = OllamaEmbeddings(model="llama3.2")

documents = []
def store_files_in_chromadb(directory):
    for root, _, files in os.walk(directory):
        for file in files:
            try:
                file_path = os.path.join(root, file)
                if ".md" in file_path:
                    print(f"Loading: {file_path}")
                    loader = UnstructuredMarkdownLoader(file_path)
                    data = loader.load()
                    for d in data:
                        d.metadata = {"source": file_path}
                        documents.append(d)
                    print("Successfully loaded")
            except Exception as e:
                print(f"Error loading file: {e}")
    rec_char_splitter = RecursiveCharacterTextSplitter(chunk_size=1000, chunk_overlap=100)
    rec_char_docs = rec_char_splitter.split_documents(documents)
    Chroma.from_documents(
        rec_char_docs, embeddings, persist_directory=persistent_directory
    )
    print("------SUCCESSFULLY STORED DOCUMENTS IN DB-----")

if not os.path.exists(persistent_directory):
   local_path = f"{current_dir}/gofr_repo/docs"
   store_files_in_chromadb(local_path)
else:
    print("------DOCUMENTS ALREADY STORED IN DB-----")



