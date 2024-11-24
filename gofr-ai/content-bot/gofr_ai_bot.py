import os
from dotenv import load_dotenv
from langchain_community.vectorstores import Chroma
from langchain_core.messages import HumanMessage, SystemMessage
from langchain_ollama import OllamaEmbeddings
from langchain_anthropic import ChatAnthropic
from fastapi import FastAPI, Request


# load_dotenv()
# model = ChatAnthropic(
#         model="claude-3-5-sonnet-20241022",
#         temperature=0,
#         max_tokens=1024,
#         timeout=None,
#         max_retries=2,
#     )


# current_dir = os.path.dirname(os.path.abspath(__file__))
# persistent_directory = os.path.join(
#     current_dir, "db", "chroma_db_llama32_claude")

# embeddings = OllamaEmbeddings(model="llama3.2")

# db = Chroma(persist_directory=persistent_directory,
#             embedding_function=embeddings)

# retriever = db.as_retriever(
#     search_type="similarity",
#     search_kwargs={"k": 5},
# )

def issueResolver(query):
    load_dotenv()
    model = ChatAnthropic(
            model="claude-3-5-sonnet-20241022",
            temperature=0.7,
            max_tokens=1024,
            timeout=None,
            max_retries=2,
        )


    current_dir = os.path.dirname(os.path.abspath(__file__))
    persistent_directory = os.path.join(
        current_dir, "db", "chroma_db_llama32_claude")

    embeddings = OllamaEmbeddings(model="llama3.2")

    db = Chroma(persist_directory=persistent_directory,
                embedding_function=embeddings)

    retriever = db.as_retriever(
        search_type="similarity",
        search_kwargs={"k": 5},
    )
    relevant_docs = retriever.invoke(query)

    combined_input = (
        "You are the maintainer of GoFr framework in Go. Based on your knowledge till date " +
        " answer the following question if GoFr framework can help in solving it\n. Question: " +
        query +
        "Note:\nProvide code examples as necessary but keep it to the point. Add links to the relevant GoFr doumentation if possible in case GoFr can solve the issue " +
        " If GoFr cannot solve it, just provide 'GoFr is not designed to solve this issue'. Don't add additional information." +
        "Here are some documents related to GoFr docs that you can refer in addition to your knowledge of GoFr framework. " + 
        "\n\nRelevant Documents:\n" +
        "\n\n".join([doc.page_content for doc in relevant_docs])
    )

    messages = [
        SystemMessage(content="You are a helpful code assistant and maintainer of a Go framework called GoFr."),
        HumanMessage(content=combined_input),
    ]

    result = model.invoke(messages)

    print("\n--- Generated Response ---")
    return result.content

def postGenerator(query):
    load_dotenv()
    model = ChatAnthropic(
            model="claude-3-5-sonnet-20241022",
            temperature=0.7,
            max_tokens=1024,
            timeout=None,
            max_retries=2,
        )


    current_dir = os.path.dirname(os.path.abspath(__file__))
    persistent_directory = os.path.join(
        current_dir, "db", "chroma_db_llama32_claude")

    embeddings = OllamaEmbeddings(model="llama3.2")

    db = Chroma(persist_directory=persistent_directory,
                embedding_function=embeddings)

    retriever = db.as_retriever(
        search_type="similarity",
        search_kwargs={"k": 5},
    )
    relevant_docs = retriever.invoke(query)

    combined_input = (
        "You are the social media marketing manager and maintainer of GoFr framework in Go. Based on your knowledge till date " +
        " generate an engaging social media post that are catered to Go developers based on the latest trends and the following query: " + 
        query +
        "Note that it is very important to consider the following:\n" +
        "1. If post:twitter is mentioned in the query, stick to the character limit of the default twitter post. If post:linkedin is mentioned in the query follow its standards and similarily for other social media platforms" +
        "2. If the query is someting that GoFr framework can help in, definitely mention how GoFr can help and specify helpful links related to the same or point them to the related GoFr doc"
        "3. If the query is someting GoFr cannot help with, just provide 'GoFr is not designed to solve this issue'. Don't add additional information"
        "4. Don't include any sentence in the begginning while starting out. Just put out the required content only."
        "Here are some documents related to GoFr docs that you can refer in addition to your knowledge of GoFr framework. " + 
        "\n\nRelevant Documents:\n" +
        "\n\n".join([doc.page_content for doc in relevant_docs])
    )

    messages = [
        SystemMessage(content="You are a skillful social media manager and maintainer of a Go framework called GoFr."),
        HumanMessage(content=combined_input),
    ]

    result = model.invoke(messages)

    print("\n--- Generated Response ---")
    return result.content

app = FastAPI( 
    title="Issue Resolver Bot GoFr",
    version="1.0",
    description= "Resolves issues related to Go and specifies how GoFr can help if applicable"
)

@app.get("/issue-resolver")
async def issue_resolve(request: Request):
    try:
        json_body = await request.json()
    except Exception:
        json_body = None
    query = json_body["content"]
    return { "result": issueResolver(query) }

@app.get("/post-generator")
async def generate_post(request: Request):
    try:
        json_body = await request.json()
    except Exception:
        json_body = None
    query = json_body["content"]
    return { "result": postGenerator( query ) }