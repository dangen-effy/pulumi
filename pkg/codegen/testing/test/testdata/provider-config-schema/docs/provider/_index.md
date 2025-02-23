
---
title: "Provider"
title_tag: "configstation.Provider"
meta_desc: "Documentation for the configstation.Provider resource with examples, input properties, output properties, lookup functions, and supporting types."
layout: api
no_edit_this_page: true
---



<!-- WARNING: this file was generated by test. -->
<!-- Do not edit by hand unless you're certain you know what you are doing! -->




## Create Provider Resource {#create}
<div>
<pulumi-chooser type="language" options="typescript,python,go,csharp,java,yaml"></pulumi-chooser>
</div>


<div>
<pulumi-choosable type="language" values="javascript,typescript">
<div class="highlight"><pre class="chroma"><code class="language-typescript" data-lang="typescript"><span class="k">new </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">name</span><span class="p">:</span> <span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p">?:</span> <span class="nx"><a href="#inputs">ProviderArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p">?:</span> <span class="nx"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span><span class="p">);</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<div class="highlight"><pre class="chroma"><code class="language-python" data-lang="python"><span class=nd>@overload</span>
<span class="k">def </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">,</span>
             <span class="nx">favorite_color</span><span class="p">:</span> <span class="nx">Optional[Union[str, Color]]</span> = None<span class="p">,</span>
             <span class="nx">secret_sandwiches</span><span class="p">:</span> <span class="nx">Optional[Sequence[_config.SandwichArgs]]</span> = None<span class="p">)</span>
<span class=nd>@overload</span>
<span class="k">def </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">resource_name</span><span class="p">:</span> <span class="nx">str</span><span class="p">,</span>
             <span class="nx">args</span><span class="p">:</span> <span class="nx"><a href="#inputs">Optional[ProviderArgs]</a></span> = None<span class="p">,</span>
             <span class="nx">opts</span><span class="p">:</span> <span class="nx"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">Optional[ResourceOptions]</a></span> = None<span class="p">)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<div class="highlight"><pre class="chroma"><code class="language-go" data-lang="go"><span class="k">func </span><span class="nx">NewProvider</span><span class="p">(</span><span class="nx">ctx</span><span class="p"> *</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span><span class="p">,</span> <span class="nx">name</span><span class="p"> </span><span class="nx">string</span><span class="p">,</span> <span class="nx">args</span><span class="p"> *</span><span class="nx"><a href="#inputs">ProviderArgs</a></span><span class="p">,</span> <span class="nx">opts</span><span class="p"> ...</span><span class="nx"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span><span class="p">) (*<span class="nx">Provider</span>, error)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">
<div class="highlight"><pre class="chroma"><code class="language-csharp" data-lang="csharp"><span class="k">public </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">string</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ProviderArgs</a></span><span class="p">? </span><span class="nx">args = null<span class="p">,</span> <span class="nx"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span><span class="p">? </span><span class="nx">opts = null<span class="p">)</span></code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<div class="highlight"><pre class="chroma">
<code class="language-java" data-lang="java"><span class="k">public </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ProviderArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">)</span>
<span class="k">public </span><span class="nx">Provider</span><span class="p">(</span><span class="nx">String</span><span class="p"> </span><span class="nx">name<span class="p">,</span> <span class="nx"><a href="#inputs">ProviderArgs</a></span><span class="p"> </span><span class="nx">args<span class="p">,</span> <span class="nx">CustomResourceOptions</span><span class="p"> </span><span class="nx">options<span class="p">)</span>
</code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<div class="highlight"><pre class="chroma"><code class="language-yaml" data-lang="yaml">type: <span class="nx">pulumi:providers:configstation</span><span class="p"></span>
<span class="p">properties</span><span class="p">: </span><span class="c">#&nbsp;The arguments to resource properties.</span>
<span class="p"></span><span class="p">options</span><span class="p">: </span><span class="c">#&nbsp;Bag of options to control resource&#39;s behavior.</span>
<span class="p"></span>
</code></pre></div>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-optional" title="Optional">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ProviderArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/nodejs/pulumi/pulumi/#CustomResourceOptions">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>resource_name</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-optional" title="Optional">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ProviderArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/python/pulumi/#pulumi.ResourceOptions">ResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">

<dl class="resources-properties"><dt
        class="property-optional" title="Optional">
        <span>ctx</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#Context">Context</a></span>
    </dt>
    <dd>Context object for the current deployment.</dd><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-optional" title="Optional">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ProviderArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="https://pkg.go.dev/github.com/pulumi/pulumi/sdk/v3/go/pulumi?tab=doc#ResourceOption">ResourceOption</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="csharp">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-optional" title="Optional">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ProviderArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>opts</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="/docs/reference/pkg/dotnet/Pulumi/Pulumi.CustomResourceOptions.html">CustomResourceOptions</a></span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">

<dl class="resources-properties"><dt
        class="property-required" title="Required">
        <span>name</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The unique name of the resource.</dd><dt
        class="property-required" title="Required">
        <span>args</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#inputs">ProviderArgs</a></span>
    </dt>
    <dd>The arguments to resource properties.</dd><dt
        class="property-optional" title="Optional">
        <span>options</span>
        <span class="property-indicator"></span>
        <span class="property-type">CustomResourceOptions</span>
    </dt>
    <dd>Bag of options to control resource&#39;s behavior.</dd></dl>

</pulumi-choosable>
</div>

## Provider Resource Properties {#properties}

To learn more about resource properties and how to use them, see [Inputs and Outputs](/docs/intro/concepts/inputs-outputs) in the Architecture and Concepts docs.

### Inputs

The Provider resource accepts the following [input](/docs/intro/concepts/inputs-outputs) properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favoritecolor_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favoritecolor_csharp" style="color: inherit; text-decoration: inherit;">Favorite<wbr>Color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string | <a href="#color">Configstation.<wbr>Pulumi.<wbr>Configstation.<wbr>Color</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secretsandwiches_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secretsandwiches_csharp" style="color: inherit; text-decoration: inherit;">Secret<wbr>Sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">List&lt;Configstation.<wbr>Pulumi.<wbr>Configstation.<wbr>Config.<wbr>Inputs.<wbr>Sandwich<wbr>Args&gt;</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favoritecolor_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favoritecolor_go" style="color: inherit; text-decoration: inherit;">Favorite<wbr>Color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string | <a href="#color">Color</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secretsandwiches_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secretsandwiches_go" style="color: inherit; text-decoration: inherit;">Secret<wbr>Sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">Sandwich<wbr>Args</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favoritecolor_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favoritecolor_java" style="color: inherit; text-decoration: inherit;">favorite<wbr>Color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String | <a href="#color">Color</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secretsandwiches_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secretsandwiches_java" style="color: inherit; text-decoration: inherit;">secret<wbr>Sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">List&lt;Sandwich<wbr>Args&gt;</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favoritecolor_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favoritecolor_nodejs" style="color: inherit; text-decoration: inherit;">favorite<wbr>Color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string | <a href="#color">Color</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secretsandwiches_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secretsandwiches_nodejs" style="color: inherit; text-decoration: inherit;">secret<wbr>Sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">config<wbr>Sandwich<wbr>Args[]</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favorite_color_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favorite_color_python" style="color: inherit; text-decoration: inherit;">favorite_<wbr>color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str | <a href="#color">Color</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secret_sandwiches_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secret_sandwiches_python" style="color: inherit; text-decoration: inherit;">secret_<wbr>sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">Sandwich<wbr>Args]</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="favoritecolor_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#favoritecolor_yaml" style="color: inherit; text-decoration: inherit;">favorite<wbr>Color</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String | <a href="#color">&#34;blue&#34; | &#34;red&#34;</a></span>
    </dt>
    <dd>this is a relaxed string enum which can also be set via env var It can also be sourced from the following environment variable: <code>FAVE_COLOR</code></dd><dt class="property-optional"
            title="Optional">
        <span id="secretsandwiches_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#secretsandwiches_yaml" style="color: inherit; text-decoration: inherit;">secret<wbr>Sandwiches</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type"><a href="#sandwich">List&lt;Property Map&gt;</a></span>
    </dt>
    <dd>Super duper secret sandwiches.</dd></dl>
</pulumi-choosable>
</div>


### Outputs

All [input](#inputs) properties are implicitly available as output properties. Additionally, the Provider resource produces the following output properties:



<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_csharp" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_go" style="color: inherit; text-decoration: inherit;">Id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_java" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_nodejs" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_python" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-"
            title="">
        <span id="id_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#id_yaml" style="color: inherit; text-decoration: inherit;">id</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd>The provider-assigned unique ID for this managed resource.</dd></dl>
</pulumi-choosable>
</div>







## Supporting Types



<h4 id="color">
Color<pulumi-choosable type="language" values="python,go" class="inline">, Color<wbr>Args</pulumi-choosable>
</h4>

<div>
<pulumi-choosable type="language" values="csharp">
<dl class="tabular"><dt>Blue</dt>
    <dd>blue</dd><dt>Red</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="tabular"><dt>Color<wbr>Blue</dt>
    <dd>blue</dd><dt>Color<wbr>Red</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="tabular"><dt>Blue</dt>
    <dd>blue</dd><dt>Red</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="nodejs">
<dl class="tabular"><dt>Blue</dt>
    <dd>blue</dd><dt>Red</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="tabular"><dt>BLUE</dt>
    <dd>blue</dd><dt>RED</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="tabular"><dt>"blue"</dt>
    <dd>blue</dd><dt>"red"</dt>
    <dd>red</dd></dl>
</pulumi-choosable>
</div>

<h4 id="sandwich">
Sandwich<pulumi-choosable type="language" values="python,go" class="inline">, Sandwich<wbr>Args</pulumi-choosable>
</h4>

<div>
<pulumi-choosable type="language" values="csharp">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_csharp" style="color: inherit; text-decoration: inherit;">Bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_csharp">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_csharp" style="color: inherit; text-decoration: inherit;">Veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;string&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="go">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_go" style="color: inherit; text-decoration: inherit;">Bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_go">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_go" style="color: inherit; text-decoration: inherit;">Veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">[]string</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="java">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_java" style="color: inherit; text-decoration: inherit;">bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_java">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_java" style="color: inherit; text-decoration: inherit;">veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="javascript,typescript">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_nodejs" style="color: inherit; text-decoration: inherit;">bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_nodejs">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_nodejs" style="color: inherit; text-decoration: inherit;">veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">string[]</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="python">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_python" style="color: inherit; text-decoration: inherit;">bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">str</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_python">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_python" style="color: inherit; text-decoration: inherit;">veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">Sequence[str]</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>

<div>
<pulumi-choosable type="language" values="yaml">
<dl class="resources-properties"><dt class="property-optional"
            title="Optional">
        <span id="bread_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#bread_yaml" style="color: inherit; text-decoration: inherit;">bread</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">String</span>
    </dt>
    <dd></dd><dt class="property-optional"
            title="Optional">
        <span id="veggies_yaml">
<a data-swiftype-name="resource-property" data-swiftype-type="text" href="#veggies_yaml" style="color: inherit; text-decoration: inherit;">veggies</a>
</span>
        <span class="property-indicator"></span>
        <span class="property-type">List&lt;String&gt;</span>
    </dt>
    <dd></dd></dl>
</pulumi-choosable>
</div>


<h2 id="package-details">Package Details</h2>
<dl class="package-details">
	<dt>Repository</dt>
	<dd><a href="">configstation </a></dd>
	<dt>License</dt>
	<dd></dd>
</dl>

