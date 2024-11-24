#!/usr/bin/python3
import imaplib
import email
from email.header import decode_header
import json

# Your email credentials
EMAIL = "argoats82@gmail.com"
PASSWORD = "rfmaillehtozdfpy"

def clean(text):
    # Clean text for creating filenames
    return "".join(c if c.isalnum() else "_" for c in text)

def gatherMail():
    emails = []
    # Connect to the server
    imap = imaplib.IMAP4_SSL("imap.gmail.com")

    # Login to your account
    imap.login(EMAIL, PASSWORD)

    # Select the mailbox you want to use
    imap.select("inbox")

    # Search for all UNSEEN emails
    status, messages = imap.search(None, 'UNSEEN')

    # Convert messages to a list of email IDs
    messages = messages[0].split()

    for mail in messages:
        # Fetch the email by ID
        res, msg = imap.fetch(mail, "(RFC822)")
        for response in msg:
            if isinstance(response, tuple):
                # Parse a bytes email into a message object
                msg = email.message_from_bytes(response[1])
                # Decode the email subject
                subject, encoding = decode_header(msg["Subject"])[0]
                if isinstance(subject, bytes):
                    # If it's a bytes type, decode to str
                    subject = subject.decode(encoding if encoding else "utf-8")
                # Decode email sender
                from_ = msg.get("From")
                
                # Initialize body as an empty string
                body = ""
                
                # If the email message is multipart
                if msg.is_multipart():
                    # Iterate over email parts
                    for part in msg.walk():
                        # Extract content type of email
                        content_type = part.get_content_type()
                        content_disposition = str(part.get("Content-Disposition"))
                        try:
                            # Get the email body
                            body = part.get_payload(decode=True).decode()
                            if content_type == "text/plain" and "attachment" not in content_disposition:
                                break  # Stop after getting the text body
                        except:
                            pass
                else:
                    # If not multipart, get the body directly
                    body = msg.get_payload(decode=True).decode()

                # Append the email details to the list
                body = body.split("--")[ 0 ]
                from_ = from_.split("<")[ 1 ][ : -1 ]
                emails.append({
                    "from": from_,
                    "subject": subject,
                    "body": body
                })
    # Close the connection and logout
    imap.close()
    imap.logout()
    output = { "output": emails }
    print( output )

gatherMail()
