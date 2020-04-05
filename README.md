# cfe

*CFE is lightness, speed and parallels library for declarative parsing without build DOM tree*

CFE is **C**ontext, **F**or-loop and **E**xtract. 

Any structing text, you can parsing very simple. We take a very simply case:

  - You need find start position of anchor (tag or tag's attribute)
  - You need find finish position of anchor (tag or tag's attribute)
  - You need go to the start anchor's position
  - Extract object until finish position anchor
  - Success!
  
## Why?

- We don't build DOM-tree for extract data. Then we have speed
- We don't use DOM-tree for extract data. Then we have very simple algorithm. We only find first substring of text
- We don't use DOM-tree for extract data. Then we can parallize any step of our algorithm in separate gouroutine
- We use only three simple operations: *Extract*, *Context* and *For-loop*
- Than we have constructions like this: For(Ctx(Extract())). This is a very simple.

## Example

Let's see at simple instance. We have such html-code:

     <x>Hello world!<z>a</z></x>
  
We want extract content of tag `<z>`:

  - You need find start position of tag `<z>`
  - You need find finish position of tag `</z>`
  - You need go to the start `<z>`'s position (16+3 (length of `<z>`) index from start since 0-index)
  - Extract object until finish `</z>`'s position  (20 index from start since 0-index)
  - Success!
  
The such way, you can extract data from attributes:

    <x>Hello world!<z data="a"></z></x>

  - You need find start position of tag `<z data=">`
  - You need find finish position of tag `">`
  - You need go to the start `<z data=">`'s position
  - Extract object until finish `">`'s position
  - Success!
  
Let's see more difficult example. We want to extract object `a`:

    <x>Hello world!<z><x>a</x></z>
    
Then we need to use the context:

  - You need find start position of tag `<z>`
 
We find the context. Next, we take the same things from the position of context:

  - You need find start position of tag `<x>`
  - You need find finish position of tag `</x>`
  - You need go to the start `<x>`'s position
  - Extract object until finish `</x>`'s position
  - Success!
  
Now, we have some such constructions. You need to use for-loop construction, while there is some objects.

That's all. It's enough.

# Type of data any sites

1. html-data
2. json-data
3. api-data
4. encrypted-api-data

We can process only 1 and 2 list items.

1 This is case have data into html:

    <div class="sku">
      <span>Код товара:</span>
      <span>123456</span>
    </div>
    
2 This is case have json object into html:

    <script type="text/javascript">
      d['H'] = {"e":"l","l":"o","w":"o","r":"l","d":"!"};
    </script>
    
3 Request to API. This case you can process with golang.
4 Request to API with encrypting data into the request. This case relally problem for us. Than, you decrypt data and after that, you can use CFE. Or, you can use another tool with headlesss browser parsing. But this is a very rarely and difficult case. 

**Notice**: while parcing take data as anchor is a very bad practice. Data can be changed. From example above, you don't use as anchor substring `Hello world!`. This is data. Data is changed more often, than tags
